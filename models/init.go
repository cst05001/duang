package models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init() {
	duangcfg, err := config.NewConfig("ini", "duang.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
    orm.RegisterDriver("mysql", orm.DR_MySQL)
    orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8", duangcfg.String("db_user"), duangcfg.String("db_pass"), duangcfg.String("db_addr"), duangcfg.String("db_name")))

    orm.RegisterModel(new(Unit))

    force := false
    verbose := true
    err = orm.RunSyncdb("default", force, verbose)
	if err != nil {
	    fmt.Println(err)
	}
}