package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
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
	ip := &core.Ip{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, ip)
	if err != nil {
		fmt.Println(err)
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	re := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+\\.\\d+$")
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

func (this *IPPoolController) List() {
	ipPool := core.NewIpPool()
	ips, err := ipPool.GetAllIP()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	WriteJson(this.Ctx, ips)
}
