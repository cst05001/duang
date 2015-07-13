package MCscheduler

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/cdevr/WapSNMP"
	"github.com/cst05001/duang/models/core"
)

type DockerdForSort struct {
	Dockerd *core.Dockerd
	Score   int64
}

type MCscheduler struct {
	DockerdForSortList []*DockerdForSort
	Community          string
	Version            wapsnmp.SNMPVersion
	Timeout            int
	Retry              int
}

/*
	本调度器算法：
		step1 内存大优先
		step2 线程数*空闲值 高者优先
*/
func NewScheduler() *MCscheduler {
	mc := &MCscheduler{}
	mc.Init()
	return mc
}

func (this *MCscheduler) Init() error {
	duangcfg, err := config.NewConfig("ini", "conf/duang.conf")
	if err != nil {
		fmt.Println(err)
		return err
	}

	v := duangcfg.String("snmp_version")
	if v == "1" || v == "v1" {
		this.Version = wapsnmp.SNMPv1
	} else if v == "2c" || v == "v2c" {
		this.Version = wapsnmp.SNMPv2c
	} else {
		err = errors.New(fmt.Sprintf("duang.cfg snmp_version should be 1 or 2c, but got: %s", v))
		beego.Error(err)
		return err
	}

	this.Community = duangcfg.String("snmp_community")
	if len(this.Community) < 1 {
		err = errors.New("duang.conf snmp_community cannot be null")
		beego.Error(err)
		return err
	}

	timeout, err := duangcfg.Int("snmp_timeout")
	if err != nil {
		beego.Error(err)
		return err
	}
	this.Timeout = timeout

	retry, err := duangcfg.Int("snmp_retry")
	if err != nil {
		beego.Error(err)
		return err
	}
	this.Retry = retry

	o := orm.NewOrm()
	o.Using("default")

	this.DockerdForSortList = make([]*DockerdForSort, 0)
	dockerdList := make([]*core.Dockerd, 0)
	_, err = o.QueryTable("dockerd").All(&dockerdList)
	if err != nil {
		beego.Error(err)
		return err
	}

	for _, dockerd := range dockerdList {
		//this.Score[dockerd] = 0
		this.DockerdForSortList = append(this.DockerdForSortList, &DockerdForSort{Dockerd: dockerd, Score: 0})
	}

	this.UpdateScore()
	return nil
}

func (this *MCscheduler) UpdateScore() error {
	//这里应该改成多线程。
	for k, _ := range this.DockerdForSortList {
		ip := this.DockerdForSortList[k].Dockerd.GetIP()
		wsnmp, err := wapsnmp.NewWapSNMP(ip, this.Community, this.Version, time.Duration(this.Timeout)*time.Second, this.Retry)
		defer wsnmp.Close()
		if err != nil {
			beego.Error("Conn: ", err)
			return err
		}

		//获取剩余内存
		v, err := wsnmp.Get(wapsnmp.MustParseOid(".1.3.6.1.4.1.2021.4.6.0"))
		if err != nil {
			beego.Error("Get SNMP .1.3.6.1.4.1.2021.4.6.0 err at: ", ip, " with error: ", err)
			return err
		}
		var memAvailReal int64
		memAvailReal = v.(int64)
		beego.Debug(ip, " memAvailReal: ", memAvailReal)

		v, err = wsnmp.Get(wapsnmp.MustParseOid(".1.3.6.1.4.1.2021.4.14.0"))
		if err != nil {
			beego.Error("Get SNMP .1.3.6.1.4.1.2021.4.14.0 err at: ", ip, " with error: ", err)
			return err
		}
		var memBuffer int64
		memBuffer = v.(int64)
		beego.Debug(ip, " memBuffer: ", memBuffer)

		v, err = wsnmp.Get(wapsnmp.MustParseOid(".1.3.6.1.4.1.2021.4.15.0"))
		if err != nil {
			beego.Error("Get SNMP .1.3.6.1.4.1.2021.4.15.0 err at: ", ip, " with error: ", err)
			return err
		}
		var memCached int64
		memCached = v.(int64)
		beego.Debug(ip, " memCached: ", memCached)

		//获取CPU信息
		//线程数
		table, err := wsnmp.GetTable(wapsnmp.MustParseOid(".1.3.6.1.2.1.25.3.3.1.2"))
		if err != nil {
			beego.Error("Get SNMP .1.3.6.1.2.1.25.3.3.1.2 err at: ", ip, " with error: ", err)
			return err
		}
		threadNum := len(table)
		beego.Debug(ip, " threadNum: ", threadNum)
		//1分钟压力
		table, err = wsnmp.GetTable(wapsnmp.MustParseOid(".1.3.6.1.4.1.2021.10.1.6.1"))
		if err != nil {
			beego.Error("Get SNMP .1.3.6.1.4.1.2021.10.1.6.1 err at: ", ip, " with error: ", err)
			return err
		}
		var load float64
		for _, v := range table {
			load = v.(float64)
			break
		}
		//计分
		//可用内存 * (线程数 * (1 - CPU负载)) = 可用内存 * (线程数 * 空闲CPU资源)
		this.DockerdForSortList[k].Score = int64(float64(memAvailReal+memBuffer+memCached) * (float64(threadNum) * (float64(1) - load)))
	}
	return nil
}

func (this *MCscheduler) GetDockerd(n int64, excludeBackends []string) []*core.Dockerd {
	this.UpdateScore()
	i := 0
	for i < len(this.DockerdForSortList) {
		j := i
		for j < len(this.DockerdForSortList) {
			if this.DockerdForSortList[j].Score > this.DockerdForSortList[i].Score {
				tmp := this.DockerdForSortList[j]
				this.DockerdForSortList[j] = this.DockerdForSortList[i]
				this.DockerdForSortList[i] = tmp
			}
			j = j + 1
		}
		i = i + 1
		if int64(i) >= n {
			break
		}
	}
	dockerdList := make([]*core.Dockerd, 0)
	i = 0
	for k, _ := range this.DockerdForSortList {
		// 如果在 exclude 列表，则跳过。
		if excludeBackends != nil {
			dockerd := dockerdList[k]
			in := false
			for _, exclude := range excludeBackends {
				if dockerd.Addr == exclude {
					in = true
					break
				}
			}
			if in == true {
				continue
			}
		}

		dockerdList = append(dockerdList, this.DockerdForSortList[k].Dockerd)
		beego.Debug("GetDockerd: ", this.DockerdForSortList[k].Dockerd.GetIP(), " :\t", this.DockerdForSortList[k].Score)
		i = i + 1
		if int64(i) >= n {
			break
		}
	}
	return dockerdList
}
