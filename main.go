package main

import (
	"github.com/astaxie/beego"
	_ "github.com/cst05001/duang/routers"
	_ "github.com/cst05001/duang/models"

)

func main() {
	beego.Run()
}

