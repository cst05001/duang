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
	beego.Router("/unit/create", &controllers.UnitController{}, "get:CreateHtml")
	beego.Router("/unit/create", &controllers.UnitController{}, "post:Create")
	beego.Router("/unit/list", &controllers.UnitController{}, "get:List")
	beego.Router("/unit/update/:unitid:int", &controllers.UnitController{}, "get:UpdateHtml")
	beego.Router("/unit/update/:unitid:int", &controllers.UnitController{}, "post:Update")
	beego.Router("/unit/run/:unitid:int", &controllers.UnitController{}, "*:Run")

	beego.Router("/dockerd/create", &controllers.DockerdController{}, "get:CreateHtml")
	beego.Router("/dockerd/create", &controllers.DockerdController{}, "post:Create")
	beego.Router("/dockerd/list", &controllers.DockerdController{}, "get:List")

	beego.Router("/ippool/create", &controllers.IPPoolController{}, "get:CreateHtml")
	beego.Router("/ippool/create", &controllers.IPPoolController{}, "post:Create")
	beego.Router("/ippool/list", &controllers.IPPoolController{}, "get:ListAll")
	beego.Router("/ippool/list/used", &controllers.IPPoolController{}, "get:ListUsed")
	beego.Router("/ippool/list/free", &controllers.IPPoolController{}, "get:ListFree")
}
