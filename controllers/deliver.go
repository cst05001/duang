package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models/core"
	"regexp"
	"strconv"
)

type DeliverController struct {
	beego.Controller
}

func (this *DeliverController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *DeliverController) FrontendCreateHtml() {
	this.TplNames = "deliver/frontend_create.tpl"
	this.Render()
}

func (this *DeliverController) FrontendCreate() {
	frontend := &core.Frontend{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, frontend)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	//re := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+\\.\\d+$")
	re := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+\\.\\d+:\\d+$")
	if !re.MatchString(frontend.Bind) {
		WriteJson(this.Ctx, &StatusError{Error: "bad format"})
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	frontend.Id, err = o.Insert(frontend)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	err = o.Read(frontend)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	WriteJson(this.Ctx, frontend)
}

func (this *DeliverController) FrontendList() {
	o := orm.NewOrm()
	o.Using("default")

	frontendList := make([]*core.Frontend, 0)
	_, err := o.QueryTable("frontend").All(&frontendList)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	WriteJson(this.Ctx, frontendList)
}

func (this *DeliverController) FrontendDel() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	frontend := &core.Frontend{Id: int64(id)}
	_, err = o.Delete(frontend)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

}
