package main

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/session"

	"gopulse/backend/models"

	_ "gopulse/backend/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlstring"))
	orm.RegisterModel(new(models.User))
	orm.Debug = true
	orm.NewOrm()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("mysqlstring")
	beego.Run()
}
