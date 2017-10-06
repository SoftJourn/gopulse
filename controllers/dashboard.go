package controllers

import (
    //"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type DashboardController struct {
	beego.Controller
}

func (this *DashboardController) Get() {
    // https://beego.me/docs/intro/

	email := this.GetSession("email")
	if email == nil {
		// Whatever
		beego.Error("Email is empty")
		this.Redirect("/", http.StatusInternalServerError) //http.StatusFound)
		return
	}
    this.Data["Name"] = this.GetSession("name")
	this.Data["Email"] = email

	//beego.Info(fmt.Sprintf("%#v", user))
	beego.Info(email)
	this.TplName = "dashboard.html"
	this.Render()
}
