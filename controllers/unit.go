package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

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

//Reviewed at 20150702
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

	//事务开始
	err = o.Begin()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	//插入unit
	unit.Id, err = o.Insert(unit)
	if err != nil {

		errRollback := o.Rollback()
		if errRollback != nil {
			fmt.Println(errRollback)
			return
		}
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	//插入参数
	for _, v := range unit.Parameteres {
		//参数不全则忽略
		if len(v.Value) == 0 || len(v.Type) == 0 {
			continue
		}

		v.Unit = unit
		_, err = o.Insert(v)
		if err != nil {
			errRollback := o.Rollback()
			if errRollback != nil {
				WriteJson(this.Ctx, &StatusError{Error: errRollback.Error()})
				return
			}
			WriteJson(this.Ctx, &StatusError{Error: err.Error()})
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
	WriteJson(this.Ctx, unit)
}

func (this *UnitController) Delete() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	unit := &core.Unit{Id: int64(id)}
	_, err = o.Delete(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
}

//Reviewed at 20150702
func (this *UnitController) List() {
	o := orm.NewOrm()
	o.Using("default")
	unitList := make([]core.Unit, 0)
	_, err := o.QueryTable("unit").All(&unitList)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	for k, _ := range unitList {
		o.LoadRelated(&unitList[k], "Parameteres")
		//下面循环是为了避免beego ORM数据结构和json.Marshal配合导致的死循环解析问题
		for _, p := range unitList[k].Parameteres {
			p.Unit = nil
		}
	}

	WriteJson(this.Ctx, unitList)
}

func (this *UnitController) UpdateHtml() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	unit := &core.Unit{Id: int64(id)}
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

//Reviewed at 20150702
func (this *UnitController) Update() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	unit := &core.Unit{Id: int64(id)}
	/* 这段代码可能没必要，因为update的时候，前端会发送全部新Unit的信息，所以不用读取旧数据
	err = o.Read(unit)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	*/

	err = o.Begin()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	// 删除旧数据全部关联参数
	_, err = o.QueryTable("UnitParameter").Filter("unit_id", int64(unit.Id)).Delete()
	if err != nil {
		errRollback := o.Rollback()
		if errRollback != nil {
			WriteJson(this.Ctx, &StatusError{Error: errRollback.Error()})
			return
		}
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	//获取Unit新值
	err = json.Unmarshal(this.Ctx.Input.RequestBody, unit)
	if err != nil {
		errRollback := o.Rollback()
		if errRollback != nil {
			WriteJson(this.Ctx, &StatusError{Error: errRollback.Error()})
			return
		}
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	unit.Id = int64(id)

	_, err = o.Update(unit)
	if err != nil {
		errRollback := o.Rollback()
		if err != nil {
			WriteJson(this.Ctx, &StatusError{Error: errRollback.Error()})
			return
		}
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
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
			errRollback := o.Rollback()
			if err != nil {
				WriteJson(this.Ctx, &StatusError{Error: errRollback.Error()})
				return
			}
			WriteJson(this.Ctx, &StatusError{Error: err.Error()})
			return
		}
	}
	err = o.Commit()
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}

	//这段代码只是为了解决json库和beego ORM死循环解析的问题
	for _, v := range unit.Parameteres {
		if len(v.Value) == 0 || len(v.Type) == 0 {
			continue
		}
		v.Unit = nil
	}

	WriteJson(this.Ctx, unit)
}

//Reviewed at 20150702
func (this *UnitController) Run() {
	unitId, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
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

	//向调度器索要指定数量的dockerd，用来运行container。调度器决定了container跑在哪几台机器上。
	unit.Dockerd = models.Scheduler.GetDockerd(unit.Number)

	//创建一个DockerEngine，载入对应引擎。DockerEngine决定了启动container的行为。
	var client dockerdengine.DockerClient
	client = dockerd_engine1.NewDockerClientEng1(unit)
	/*
		运行Unit，并附上回调函数，这是容器Create 和 Run成功、失败一共4个状态的回调函数。
		详情请参考 models/dockerengine/dockerclient.go
	*/
	err = client.Run(unit, dockerdCallbackFunc)
	if err != nil {
		WriteJson(this.Ctx, &StatusError{Error: err.Error()})
		return
	}
	/*
		由于启动docker将改成异步，所以要留一个查询Run状态的接口。
	*/
}

//Reviewed at 20150702
func dockerdCallbackFunc(dockerd *core.Dockerd, status int, args ...interface{}) {
	ippool := core.NewIpPool()
	var ip *core.Ip
	var err error
	switch status {
	case dockerdengine.STATUS_ON_CREATE_SUCCESSED:

	case dockerdengine.STATUS_ON_CREATE_FAILED:

	case dockerdengine.STATUS_ON_RUN_SUCCESSED:
		/*
			参考: models/dockerengine/engine1/container.go
			DockerEngine给回调函数传入的第一个参数，是*core.Unit。这个行为取决于各个Engine的实现，请大家遵守。
			如果不遵守，需要修改本回调函数逻辑。

			代码逻辑：Run container -> 分配IP -> 把IP提交到前端分发器
		*/
		unit := args[0].(*core.Unit)
		/*
			从ip池获取可以分配给container的IP。
			注意，这里不是dockerd宿主机的IP。是container的IP。
			因为container采用独立IP桥接模式。
			尚未做NAT网络支持。我也觉得独立桥接比NAT好。
		*/
		ip, err = ippool.GetFreeIP()
		if err != nil {
			beego.Error("Cannot get free IP at ", dockerd.GetIP(), " :", err)
			return
		}

		//通过 ssh client 调用pipework
		duangcfg, err := config.NewConfig("ini", "conf/duang.conf")
		if err != nil {
			beego.Error(err)
			ippool.ReleaseIP(ip.Id)
			return
		}

		var sshclient sshclientengine.SshClientInterface
		//通过密钥对访问ssh服务器，也就是dockerd所在的服务器，也就是宿主机。
		sshclient, err = sshclientengine1.NewSSLClient(fmt.Sprintf("%s:%s", dockerd.GetIP(), duangcfg.String("ssh_port")), duangcfg.String("ssh_user"), duangcfg.String("ssh_keypath"))
		if err != nil {
			ippool.ReleaseIP(ip.Id)
			return
		}
		//pipework br0 containerName 192.168.0.0/24@192.168.0.1
		cmd := fmt.Sprintf("%s %s %s %s", duangcfg.String("pipework_path"), duangcfg.String("pipework_bridge"), unit.Name, ip.Ip)
		err = sshclient.Run(cmd)
		if err != nil {
			ippool.ReleaseIP(ip.Id)
			return
		}

		//调用前端分发器，把container IP分发到 confd+HAProxy
		var deliverEngine deliverengine.DeliverInterface
		deliverEngine = deliver_engine1.NewDeliver()
		//解析出Unit参数中的前端分发器参数。
		for _, para := range unit.Parameteres {
			if para.Type == "d" {
				re := regexp.MustCompile("^(\\d+):(\\d+)$")
				result := re.FindStringSubmatch(para.Value)
				fPort := result[1]
				bPort := result[2]

				err = deliverEngine.Bind(fmt.Sprintf("0.0.0.0:%s", fPort), unit.Domain, fmt.Sprintf("%s:%s", ip.GetIP(), bPort))
				if err != nil {
					beego.Error(err)
				}
			}
		}

	case dockerdengine.STATUS_ON_RUN_FAILED:
	}
}
