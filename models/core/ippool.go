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
func (this *IpPool) GetFreeIP(n int64) ([]*Ip, error) {
	ips, err := this.ListIPByStatus(1, n)
	if err != nil {
		return nil, err
	}

	err = this.ChangeState(ips, 0)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func (this *IpPool) ListUsedIP(n int64) ([]*Ip, error) {
	ips, err := this.ListIPByStatus(0, n)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func (this *IpPool) ListFreeIP(n int64) ([]*Ip, error) {
	ips, err := this.ListIPByStatus(1, n)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func (this *IpPool) ChangeState(iplist []*Ip, state uint8) error {
	o := orm.NewOrm()
	o.Using("default")

	err := o.Begin()
	if err != nil {
		return err
	}
	for _, ip := range iplist {
		ip.Status = 0
		_, err = o.Update(ip)
		if err != nil {
			fmt.Println(err)
			err = o.Rollback()
			if err != nil {
				return err
			}
			return err
		}
	}
	err = o.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (this *IpPool) ListIPById(n int64) ([]*Ip, error) {
	return nil, nil
}

//获取所有IP
func (this *IpPool) ListAllIP() ([]*Ip, error) {
	o := orm.NewOrm()
	o.Using("default")

	ips := make([]*Ip, 0)
	_, err := o.QueryTable("Ip").All(&ips)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

func (this *IpPool) ListIPByStatus(status uint8, n int64) ([]*Ip, error) {
	o := orm.NewOrm()
	o.Using("default")
	var err error

	cnt, err := o.QueryTable("Ip").Filter("status", status).Count()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ips := make([]*Ip, 0)
	if n <= cnt {
		if n < 1 {
			_, err = o.QueryTable("Ip").Filter("status", status).All(&ips)
		} else {
			_, err = o.QueryTable("Ip").Filter("status", status).Limit(n).All(&ips)
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		if n < 1 {
			_, err = o.QueryTable("Ip").Filter("status", status).All(&ips)
		} else {
			_, err = o.QueryTable("Ip").Filter("status", status).Limit(cnt).All(&ips)
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return ips, nil
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
