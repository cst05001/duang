package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models"
	engine "github.com/cst05001/duang/models/dockerclienteng1"
	"strconv"
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
		if len(v.Value) == 0 || len(v.Type) == 0 {
			continue
		}
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

func (this *UnitController) UpdateHtml() {
	unitId, err := strconv.Atoi(this.Ctx.Input.Param(":unitid"))
	if err != nil {
		fmt.Println(err)
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	unit := &models.Unit{Id: int64(unitId)}
	err = o.Read(unit)
	if err != nil {
		fmt.Println(err)
		return
	}

	o.LoadRelated(unit, "Parameteres")
	this.Data["Unit"] = unit
	this.TplNames = "unit/update.tpl"
	this.Render()
}

func (this *UnitController) Update() {
	unitId, err := strconv.Atoi(this.Ctx.Input.Param(":unitid"))
	if err != nil {
		fmt.Println(err)
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	unit := &models.Unit{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, unit)
	if err != nil {
		fmt.Println(err)
		return
	}
	unit.Id = int64(unitId)

	// 此处应支持事物，还没实现，是个 bug 要注意。
	_, err = o.Update(unit)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = o.QueryTable("UnitParameter").Filter("unit_id", int64(unit.Id)).Delete()

	for _, v := range unit.Parameteres {
		if len(v.Value) == 0 || len(v.Type) == 0 {
			continue
		}
		v.Unit = unit
		_, err = o.Insert(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	// bug 结束
	fmt.Println(unit)
	this.Ctx.WriteString("{\"status\": \"success\"}")
}

func (this *UnitController) Run() {
	unitId, err := strconv.Atoi(this.Ctx.Input.Param(":unitid"))
	if err != nil {
		fmt.Println(err)
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	unit := &models.Unit{Id: int64(unitId)}
	err = o.Read(unit)
	if err != nil {
		fmt.Println(err)
		return
	}

	o.LoadRelated(unit, "Parameteres")

	client := engine.NewDockerClient("tcp://192.168.119.10:2375")
	/*
		containerCreateResponse := client.CreateContainer(unit)
		fmt.Printf("CreateContainer: %v\n", containerCreateResponse)
		err = client.StartContainer(containerCreateResponse.ID, unit)
	*/
	err = client.Run(unit)
	if err != nil {
		fmt.Printf("Start Container Failed: %s\n", err)
		return
	}
	fmt.Println(unit)
}
