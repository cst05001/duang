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
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
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
	WriteJson(this.Ctx, unit)
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

	WriteJson(this.Ctx, unitList)
}

func (this *UnitController) UpdateHtml() {
	unitId, err := strconv.Atoi(this.Ctx.Input.Param(":unitid"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	unit := &core.Unit{Id: int64(unitId)}
	err = o.Read(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
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
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	unit := &core.Unit{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	unit.Id = int64(unitId)

	// 事务开始
	err = o.Begin()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	_, err = o.Update(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
		err = o.Rollback()
		if err != nil {
			WriteJson(this.Ctx, &StatusError{Error: err.Error()})
			return
		}
		return
	}

	// 删除全部关联参数
	_, err = o.QueryTable("UnitParameter").Filter("unit_id", int64(unit.Id)).Delete()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
		err = o.Rollback()
		if err != nil {
			WriteJson(this.Ctx, &StatusError{Error: err.Error()})
			return
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
			WriteJson(this.Ctx, &StatusError{Error: err.Error()})
			return
			err = o.Rollback()
			if err != nil {
				WriteJson(this.Ctx, &StatusError{Error: err.Error()})
				return
			}
		}
	}

	err = o.Commit()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		o.Rollback()
		return
	}
	fmt.Println(unit)
	WriteJson(this.Ctx, unit)
}

func (this *UnitController) Run() {
	unitId, err := strconv.Atoi(this.Ctx.Input.Param(":unitid"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	unit := &core.Unit{Id: int64(unitId)}
	err = o.Read(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
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
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	fmt.Println(unit)
}
