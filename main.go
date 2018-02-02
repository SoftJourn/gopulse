package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"

	"gopulse/controllers"
    "gopulse/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlstring"))
	orm.RegisterModel(new(models.User))
	orm.Debug = true
    orm.NewOrm() 
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("mysqlstring")

	beego.Router("/", &controllers.LoginController{})
	beego.Router("/auth/:id:string/login", &controllers.LoginController{}, "get:LoginHandler")
	beego.Router("/auth/:id:string/callback", &controllers.LoginController{}, "get:CallbackHandler")
	beego.Router("/dashboard/", &controllers.DashboardController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	beego.Run()
}
