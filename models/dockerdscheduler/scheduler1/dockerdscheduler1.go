package scheduler1

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models/core"
)

type DockerdScheduler1 struct {
	Count int64
}

func NewDockerdScheduler1() *DockerdScheduler1 {
	return &DockerdScheduler1{Count: 0}
}

func (this *DockerdScheduler1) GetDockerd(n int64) []*core.Dockerd {
	o := orm.NewOrm()
	o.Using("default")

	cnt, err := o.QueryTable("dockerd").Count()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if n > cnt {
		fmt.Printf("DockerdScheduler1: Dockerd needed more than dockerd has.\n")
		n = cnt
	}

	dockerdList := make([]*core.Dockerd, 0)
	_, err = o.QueryTable("dockerd").Limit(n, this.Count).All(&dockerdList)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if this.Count + n > cnt {
		dockerdList2 := make([]*core.Dockerd, 0)
		_, err = o.QueryTable("dockerd").Limit(this.Count + n - cnt, 0).All(&dockerdList)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		dockerdList = append(dockerdList, dockerdList2...)
	}
	this.Count = (this.Count + n) % cnt
	return dockerdList
}
