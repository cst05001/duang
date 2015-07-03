package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models/core"
	"regexp"
	"strconv"
	//engine "github.com/cst05001/duang/models/dockerclienteng1"
)

type DockerdController struct {
	beego.Controller
}

func (this *DockerdController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *DockerdController) Delete() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	dockerd := &core.Dockerd{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, dockerd)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	dockerd.Id = int64(id)
	o := orm.NewOrm()
	o.Using("default")
	_, err = o.Delete(dockerd)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
}

func (this *DockerdController) Update() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	dockerd := &core.Dockerd{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, dockerd)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	dockerd.Id = int64(id)
	o := orm.NewOrm()
	o.Using("default")
	_, err = o.Update(dockerd)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
}

func (this *DockerdController) CreateHtml() {
	this.TplNames = "dockerd/create.tpl"
	this.Render()
}

func (this *DockerdController) Create() {
	o := orm.NewOrm()
	o.Using("default")

	dockerd := &core.Dockerd{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, dockerd)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	reAddr := regexp.MustCompile("(.+)://(.+):(.+)")
	if !reAddr.MatchString(dockerd.Addr) {
		WriteJson(this.Ctx, &StatusError{Error: "bad format"})
		return
	}
	dockerd.Id, err = o.Insert(dockerd)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	fmt.Println(dockerd)
	WriteJson(this.Ctx, dockerd)
}

func (this *DockerdController) List() {
	o := orm.NewOrm()
	o.Using("default")
	dockerdList := make([]core.Dockerd, 0)
	_, err := o.QueryTable("dockerd").All(&dockerdList)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	for k, _ := range dockerdList {
		o.LoadRelated(&dockerdList[k], "Unit")
	}

	WriteJson(this.Ctx, dockerdList)
}
