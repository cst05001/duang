package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models"
)

type DockerdController struct {
	beego.Controller
}

func (this *DockerdController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *DockerdController) CreateHtml() {
	this.TplNames = "dockerd/create.tpl"
	this.Render()
}

func (this *DockerdController) Create() {
	o := orm.NewOrm()
	o.Using("default")

	dockerd := &models.Dockerd{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, dockerd)
	if err != nil {
		fmt.Println(err)
		return
	}

	dockerd.Id, err = o.Insert(dockerd)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dockerd)
	this.Ctx.WriteString("{\"status\": \"success\"}")
}

func (this *DockerdController) List() {
	o := orm.NewOrm()
	o.Using("default")
	dockerdList := make([]models.Dockerd, 0)
	_, err := o.QueryTable("dockerd").All(&dockerdList)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, _ := range dockerdList {
		o.LoadRelated(&dockerdList[k], "Unit")
	}

	this.Data["DockerdList"] = dockerdList
	this.TplNames = "dockerd/list.tpl"
	this.Render()
}
