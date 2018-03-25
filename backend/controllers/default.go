package controllers

import "github.com/astaxie/beego"

//	"io"

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Home() {
	c.TplName = "layouts/home.tpl"
}

func (c *MainController) Prepare() {
	c.Layout = "layouts/application.tpl"
}
