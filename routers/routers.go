package routers

import (
	"gopulse/backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.DelStaticPath("/static")
	beego.SetStaticPath("/assets", "public/assets")
	beego.SetStaticPath("/www", "./app/www")
	beego.Router("*", &controllers.MainController{}, "get:Home")
	//beego.Router("/session/login", &controllers.SessionController{}, "get:New")
	//beego.Router("/session/login", &controllers.SessionController{}, "post:Create")
	//beego.Router("/session/logout", &controllers.SessionController{}, "get:Destroy")
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/auth/:id:string/login", &controllers.LoginController{}, "get:LoginHandler")
	beego.Router("/auth/:id:string/callback", &controllers.LoginController{}, "get:CallbackHandler")
	beego.Router("/dashboard/", &controllers.DashboardController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	addGraphQLHandler()
	//addFilters()
}

func addGraphQLHandler() {
	beego.Router("/graphql", &controllers.GraphqlController{}, "post:Query")
	beego.Router("/graphql", &controllers.GraphqlController{}, "get:Index")
}

func addFilters() {
	//beego.InsertFilter("/*", beego.BeforeRouter, filterLoggedInUser)
}
