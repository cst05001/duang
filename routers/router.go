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
}
