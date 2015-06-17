package routers

import (
	"github.com/cst05001/duang/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/unit/create", &controllers.UnitController{}, "get:CreateHtml")
    beego.Router("/unit/create", &controllers.UnitController{}, "post:Create")
}
