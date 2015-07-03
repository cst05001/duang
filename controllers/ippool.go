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

func (this *IPPoolController) Delete() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	o := orm.NewOrm()
	o.Using("default")
	ip := &core.Ip{Id: int64(id)}
	_, err = o.Delete(ip)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
}

func (this *IPPoolController) Create() {
	ip := &core.Ip{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, ip)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	//re := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+\\.\\d+$")
	re := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+\\.\\d+/\\d+@\\d+\\.\\d+\\.\\d+\\.\\d+$")
	if !re.MatchString(ip.Ip) {
		WriteJson(this.Ctx, &StatusError{Error: "bad format"})
		return
	}

	ipPool := core.NewIpPool()
	ip, err = ipPool.AddIP(ip)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	WriteJson(this.Ctx, ip)
}

func (this *IPPoolController) ListAll() {

	ipPool := core.NewIpPool()
	ips, err := ipPool.ListAllIP()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	WriteJson(this.Ctx, ips)
}

func (this *IPPoolController) Release() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	ipPool := core.NewIpPool()
	err = ipPool.ReleaseIP(int64(id))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
}

func (this *IPPoolController) ListUsed() {
	ipPool := core.NewIpPool()
	ips, err := ipPool.ListUsedIP(-1)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	WriteJson(this.Ctx, ips)
}

func (this *IPPoolController) ListFree() {
	ipPool := core.NewIpPool()
	ips, err := ipPool.ListFreeIP(-1)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	WriteJson(this.Ctx, ips)
}
