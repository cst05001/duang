package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/duang/models"
	"github.com/cst05001/duang/models/core"
	"github.com/cst05001/duang/models/deliverengine"
	deliver_engine1 "github.com/cst05001/duang/models/deliverengine/engine1"
	"github.com/cst05001/duang/models/dockerdengine"
	dockerd_engine1 "github.com/cst05001/duang/models/dockerdengine/engine1"
	"github.com/cst05001/duang/models/sshclientengine"
	sshclientengine1 "github.com/cst05001/duang/models/sshclientengine/engine1"
	"regexp"
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

	// 事务开始
	err = o.Begin()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	unit.Id, err = o.Insert(unit)
	if err != nil {
		fmt.Println(err)
		err = o.Rollback()
		if err != nil {
			fmt.Println(err)
		}
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
			err = o.Rollback()
			if err != nil {
				fmt.Println(err)
			}
			return
		}

	}
	o.Commit()

	err = o.Read(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o.LoadRelated(unit, "Parameteres")
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
		for _, p := range unitList[k].Parameteres {
			p.Unit = nil
		}
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
	unit.Id = int64(unitId)
	err = o.Read(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	//o.LoadRelated(unit, "Parameteres")

	err = o.Begin()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	// 删除全部关联参数
	_, err = o.QueryTable("UnitParameter").Filter("unit_id", int64(unit.Id)).Delete()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		err = o.Rollback()
		if err != nil {
			WriteJson(this.Ctx, &StatusError{Error: err.Error()})
			return
		}
		return
	}

	err = json.Unmarshal(this.Ctx.Input.RequestBody, unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	unit.Id = int64(unitId)

	_, err = o.Update(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		err = o.Rollback()
		if err != nil {
			WriteJson(this.Ctx, &StatusError{Error: err.Error()})
			return
		}

		return
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
			err = o.Rollback()
			if err != nil {
				WriteJson(this.Ctx, &StatusError{Error: err.Error()})
				return
			}

			return
		}
	}
	err = o.Commit()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	fmt.Println(unit)
	for _, v := range unit.Parameteres {
		if len(v.Value) == 0 || len(v.Type) == 0 {
			continue
		}
		v.Unit = nil
	}

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

	var client dockerdengine.DockerClient
	client = dockerd_engine1.NewDockerClientEng1(unit)
	err = client.Run(unit, dockerdCallbackFunc)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	fmt.Println(unit)
}

func dockerdCallbackFunc(dockerd *core.Dockerd, status int, args ...interface{}) {
	ippool := core.NewIpPool()
	var ip *core.Ip
	var err error
	switch status {
	case dockerdengine.STATUS_ON_CREATE_SUCCESSED:
		fmt.Printf("在宿主机 %s 创建容器成功\n", dockerd.GetIP())

	case dockerdengine.STATUS_ON_CREATE_FAILED:
		fmt.Printf("CreateFailed: %s\n", dockerd.GetIP())
	case dockerdengine.STATUS_ON_RUN_SUCCESSED:
		unit := args[0].(*core.Unit)
		fmt.Printf("RunSuccessed: %s\n", dockerd.GetIP())

		//分配容器IP开始
		ip, err = ippool.GetFreeIP()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("分配容器IP: %s\n", ip.GetIP())
		//分配容器IP结束

		//调用pipework开始
		duangcfg, err := config.NewConfig("ini", "conf/duang.conf")
		if err != nil {
			fmt.Println(err)
			return
		}

		var sshclient sshclientengine.SshClientInterface
		sshclient, err = sshclientengine1.NewSSLClient(fmt.Sprintf("%s:%s", dockerd.GetIP(), duangcfg.String("ssh_port")), duangcfg.String("ssh_user"), duangcfg.String("ssh_keypath"))
		if err != nil {
			fmt.Println(err)
			ippool.ReleaseIP(ip.Id)
			return
		}
		//pipework br0 containerName 192.168.0.0/24@192.168.0.1
		cmd := fmt.Sprintf("%s %s %s %s", duangcfg.String("pipework_path"), duangcfg.String("pipework_bridge"), unit.Name, ip.Ip)
		fmt.Printf("CMD: %s\n", cmd)
		err = sshclient.Run(cmd)
		if err != nil {
			fmt.Println(err)
			ippool.ReleaseIP(ip.Id)
		}
		//调用pipework结束
		var deliverEngine deliverengine.DeliverInterface
		deliverEngine = deliver_engine1.NewDeliver()
		for _, para := range unit.Parameteres {
			if para.Type == "d" {
				re := regexp.MustCompile("^(\\d+):(\\d+)$")
				result := re.FindStringSubmatch(para.Value)
				fPort := result[1]
				bPort := result[2]

				err = deliverEngine.Bind(fmt.Sprintf("0.0.0.0:%s", fPort), unit.Domain, fmt.Sprintf("%s:%s", ip.GetIP(), bPort))
				if err != nil {
					fmt.Println(err)
				}
			}
		}

	case dockerdengine.STATUS_ON_RUN_FAILED:
		fmt.Printf("RunFailed: %s\n", dockerd.GetIP())
	}
}
