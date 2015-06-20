package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models"
)

type UnitController struct {
	beego.Controller
}

func (this *UnitController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *UnitController) CreateHtml() {
	this.TplNames = "unit/create.tpl"
	this.Render()
}

func (this *UnitController) Create() {
	o := orm.NewOrm()
	o.Using("default")

	unit := &models.Unit{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, unit)
	if err != nil {
		fmt.Println(err)
		return
	}

	unit.Id, err = o.Insert(unit)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range unit.Parameteres {
		v.Unit = unit
		_, err = o.Insert(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(unit)
	this.Ctx.WriteString("{\"status\": \"success\"}")
}

func (this *UnitController) List() {
	o := orm.NewOrm()
	o.Using("default")
	unitList := make([]models.Unit, 0)
	_, err := o.QueryTable("unit").All(&unitList)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, _ := range unitList {
		o.LoadRelated(&unitList[k], "Parameteres")
	}

	this.Data["UnitList"] = unitList
	this.TplNames = "unit/list.tpl"
	this.Render()
}
