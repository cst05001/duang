package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models/core"
	"regexp"
)

type IPPoolController struct {
	beego.Controller
}

func (this *IPPoolController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *IPPoolController) CreateHtml() {
	this.TplNames = "ippool/create.tpl"
	this.Render()
}

func (this *IPPoolController) Create() {
	o := orm.NewOrm()
	o.Using("default")

	ipPool := &core.IpPool{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, ipPool)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	fmt.Println(ipPool)

	reAddr := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+\\.\\d+$")
	if !reAddr.MatchString(ipPool.IP) {
		WriteJson(this.Ctx, &StatusError{Error: "bad format"})
		return
	}
	ipPool.Id, err = o.Insert(ipPool)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	fmt.Println(ipPool)
	WriteJson(this.Ctx, ipPool)
}
