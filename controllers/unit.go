package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models"
	"github.com/cst05001/duang/models/core"
	engine "github.com/cst05001/duang/models/dockerdengine/engine1"
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

	unit := &core.Unit{}
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
	unitList := make([]core.Unit, 0)
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
	unit := &core.Unit{Id: int64(unitId)}
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

	unit := &core.Unit{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, unit)
	if err != nil {
		fmt.Println(err)
		return
	}
	unit.Id = int64(unitId)

	// 事务开始
	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = o.Update(unit)
	if err != nil {
		fmt.Println(err)
		err = o.Rollback()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// 删除全部关联参数
	_, err = o.QueryTable("UnitParameter").Filter("unit_id", int64(unit.Id)).Delete()
	if err != nil {
		fmt.Println(err)
		err = o.Rollback()
		if err != nil {
			fmt.Println(err)
		}
	}

	// 插入新的参数
	for _, v := range unit.Parameteres {
		if len(v.Value) == 0 || len(v.Type) == 0 {
			continue
		}
		v.Unit = unit
		_, err = o.Insert(v)
		if err != nil {
			fmt.Println(err)
			err = o.Rollback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	err = o.Commit()
	if err != nil {
		fmt.Println(err)
	}
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
	unit := &core.Unit{Id: int64(unitId)}
	err = o.Read(unit)
	if err != nil {
		fmt.Println(err)
		return
	}

	o.LoadRelated(unit, "Parameteres")

	unit.Dockerd = models.Scheduler.GetDockerd(unit.Number)

	var client core.DockerClient
	client = engine.NewDockerClientEng1(unit)
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
