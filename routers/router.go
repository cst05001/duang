package routers

import (
	"github.com/astaxie/beego"
	"github.com/cst05001/duang/controllers"
)

func init() {
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/css", "static/css")

	beego.Router("/", &controllers.MainController{})

	beego.Router("/unit/list", &controllers.UnitController{}, "get:List")
	beego.Router("/unit/create", &controllers.UnitController{}, "get:CreateHtml")
	beego.Router("/unit/create", &controllers.UnitController{}, "post:Create")
	beego.Router("/unit/:id:int/update", &controllers.UnitController{}, "get:UpdateHtml")
	beego.Router("/unit/:id:int/update", &controllers.UnitController{}, "post:Update")
	beego.Router("/unit/:id:int/start", &controllers.UnitController{}, "get:Start")
	beego.Router("/unit/:id:int/delete", &controllers.UnitController{}, "get:Delete")
	beego.Router("/unit/:id:int/containers", &controllers.UnitController{}, "get:Containers")
	beego.Router("/unit/:id:int/status", &controllers.UnitController{}, "get:Status")
	beego.Router("/unit/:id:int/stop", &controllers.UnitController{}, "get:Stop")

	beego.Router("/dockerd/list", &controllers.DockerdController{}, "get:List")
	beego.Router("/dockerd/create", &controllers.DockerdController{}, "get:CreateHtml")
	beego.Router("/dockerd/create", &controllers.DockerdController{}, "post:Create")
	beego.Router("/dockerd/:id:int/update", &controllers.DockerdController{}, "get:Update")
	beego.Router("/dockerd/:id:int/delete", &controllers.DockerdController{}, "get:Delete")

	beego.Router("/ippool/list", &controllers.IPPoolController{}, "get:ListAll")
	beego.Router("/ippool/list/used", &controllers.IPPoolController{}, "get:ListUsed")
	beego.Router("/ippool/list/free", &controllers.IPPoolController{}, "get:ListFree")
	beego.Router("/ippool/create", &controllers.IPPoolController{}, "get:CreateHtml")
	beego.Router("/ippool/create", &controllers.IPPoolController{}, "post:Create")
	beego.Router("/ippool/:id:int/release", &controllers.IPPoolController{}, "get:Release")
	beego.Router("/ippool/:id:int/delete", &controllers.IPPoolController{}, "get:Delete")

}
