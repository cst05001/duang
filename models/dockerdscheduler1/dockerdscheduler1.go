package dockerdscheduler1

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models"
)

type DockerdScheduler1 struct {
}

func NewDockerdScheduler1() *DockerdScheduler1 {
	return &DockerdScheduler1{}
}

func (this *DockerdScheduler1) GetDockerd(n int64) []*models.Dockerd {
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

	dockerdList := make([]*models.Dockerd, 0)
	_, err = o.QueryTable("dockerd").Limit(n).All(&dockerdList)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dockerdList
}
