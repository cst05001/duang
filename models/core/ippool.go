package core

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

/*
	此功能仅限用bridge独立IP方式运行container使用
*/
type IpPool struct {
}

type Ip struct {
	Id     int64
	Ip     string `orm:"unique"`
	Status uint8  `orm:"default(1)"`
}

func NewIpPool() *IpPool {
	ipPool := &IpPool{}
	return ipPool
}

//获取可用IP
func (this *IpPool) GetFreeIP(n int) []Ip {
	return nil
}

func (this *IpPool) GetIPById(n int64) ([]*Ip, error) {
	o := orm.NewOrm()
	o.Using("default")

	cnt, err := o.QueryTable("Ip").Filter("status", 1).Count()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ips := make([]*Ip, 0)
	if n <= cnt {
		_, err := o.QueryTable("Ip").Filter("status", 1).Limit(n).All(&ips)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		_, err = o.QueryTable("Ip").Filter("status", 1).Limit(cnt).All(&ips)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	err = o.Begin()
	if err != nil {
		return nil, err
	}
	for _, ip := range ips {
		ip.Status = 0
		_, err = o.Update(ip)
		if err != nil {
			fmt.Println(err)
			err = o.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
	}
	err = o.Commit()
	if err != nil {
		return nil, err
	}

	return ips, nil
}

//获取所有IP
func (this *IpPool) GetAllIP() ([]*Ip, error) {
	o := orm.NewOrm()
	o.Using("default")

	ips := make([]*Ip, 0)
	_, err := o.QueryTable("Ip").All(&ips)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

//获取被占用的IP
func (this *IpPool) GetUsedIP() []string {
	return nil
}

//把新IP加入池
func (this *IpPool) AddIP(ip *Ip) (*Ip, error) {
	o := orm.NewOrm()
	o.Using("default")

	ip.Status = 1

	var err error
	ip.Id, err = o.Insert(ip)
	if err != nil {
		return nil, err
	}

	err = o.Read(ip)
	if err != nil {
		return nil, err
	}

	return ip, nil
}

//从IP池删除IP
func (this *IpPool) DelIP(ip string) error {
	return nil
}
