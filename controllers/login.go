package controllers

import (
	"fmt"
	"gopulse/backend/models"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Login by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Login
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LoginController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Login
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Login
// @Failure 403
// @router / [get]
func (c *LoginController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Login
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 200 {object} models.Login
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LoginController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Login
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LoginController) Delete() {

}

func (this *LoginController) Get() {

	this.TplName = "login.html"
	this.Render()
}

func (this *LoginController) LoginHandler() {
	providerName := this.Ctx.Input.Param(":id")
	beego.Info("Provider name is ", providerName)

	// setup the providers
	gomniauth.SetSecurityKey(beego.AppConfig.String("securitykey")) // NOTE: DO NOT COPY THIS - MAKE YOR OWN!
	gomniauth.WithProviders(
		github.New(beego.AppConfig.String("githubid"), beego.AppConfig.String("githubkey"), beego.AppConfig.String("githubcallback")),
		google.New(beego.AppConfig.String("googleid"), beego.AppConfig.String("googlekey"), beego.AppConfig.String("googlecallback")),
		facebook.New(beego.AppConfig.String("facebookid"), beego.AppConfig.String("facebookkey"), beego.AppConfig.String("facebookcallback")),
	)

	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		panic(err)
	}

	state := gomniauth.NewState("after", "success")

	// This code borrowed from goweb example and not fixed.
	// if you want to request additional scopes from the provider,
	// pass them as login?scope=scope1,scope2
	//options := objx.MSI("scope", ctx.QueryValue("scope"))

	authUrl, err := provider.GetBeginAuthURL(state, nil)

	if err != nil {
		this.Redirect("/", http.StatusInternalServerError)
		return
	}

	this.Redirect(authUrl, http.StatusFound)
}

func (this *LoginController) CallbackHandler() {
	providerName := this.Ctx.Input.Param(":id")
	beego.Info("Provider name is ", providerName)

	// setup the providers
	gomniauth.SetSecurityKey(beego.AppConfig.String("securitykey")) // NOTE: DO NOT COPY THIS - MAKE YOR OWN!
	gomniauth.WithProviders(
		github.New(beego.AppConfig.String("githubid"), beego.AppConfig.String("githubkey"), beego.AppConfig.String("githubcallback")),
		google.New(beego.AppConfig.String("googleid"), beego.AppConfig.String("googlekey"), beego.AppConfig.String("googlecallback")),
		facebook.New(beego.AppConfig.String("facebookid"), beego.AppConfig.String("facebookkey"), beego.AppConfig.String("facebookcallback")),
	)

	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		panic(err)
	}

	omap, err := objx.FromURLQuery(this.Ctx.Input.URI())
	if err != nil {
		beego.Error(err.Error())
		this.Redirect("/", http.StatusInternalServerError)
		return
	}

	creds, err := provider.CompleteAuth(omap)

	if err != nil {
		beego.Error(err.Error())
		this.Redirect("/", http.StatusInternalServerError)
		return
	}

	/*
		// This code borrowed from goweb example and not fixed.
		// get the state
		state, err := gomniauth.StateFromParam(ctx.QueryValue("state"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// redirect to the 'after' URL
		afterUrl := state.GetStringOrDefault("after", "error?e=No after parameter was set in the state")

	*/

	// load the user
	user, userErr := provider.GetUser(creds)

	if userErr != nil {
		beego.Error(userErr.Error())
		this.Redirect("/", http.StatusInternalServerError)
		return
	}

	t, err := models.CreateProviderUser(user, provider.Name())
	if err != nil {
		beego.Error(err.Error())
		this.Redirect("/", http.StatusInternalServerError)
		return
	}
	models.DefaultUserManager.Save(t)

	/*
	   &{map[
	     family_name:Okhotnikov%20
	     gender:male%20
	     locale:uk%20
	     hd:softjourn.com%20
	     creds:map[
	       google:%!s(*common.Credentials=&{map[
	         access_token:ya29.GlyrBAc_56HDoJusejlx8JKUvcZgwHS1KG2uhJ2sPSc3SufSJnm36_vT9koZ28daZjxNJ3s5_2wtzORplhazLg4cOyBMUvBN4ttMyDE_477Je_NCuZCw58SLsdb8Ag%20expires_in:3599000000000%20id_token:eyJhbGciOiJSUzI1NiIsImtpZCI6IjczNmU1ZjE5MTFkOTUwMTExYTU4YTkxYmJjNzUwYmRkMDgyOTk4ZmMifQ.eyJhenAiOiIxMDUxNzA5Mjk2Nzc4LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiMTA1MTcwOTI5Njc3OC5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjEwMTA4NTEyMzUzODk0MDIwNjUwMyIsImhkIjoic29mdGpvdXJuLmNvbSIsImVtYWlsIjoiYW9raG90bmlrb3ZAc29mdGpvdXJuLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhdF9oYXNoIjoiaWtuNzVFWlJ2SGdBTV9sSU1qREFLQSIsImlzcyI6ImFjY291bnRzLmdvb2dsZS5jb20iLCJpYXQiOjE1MDMwNTMxNjAsImV4cCI6MTUwMzA1Njc2MH0.JfwkywIREWRuEiwYXnbh6I14fyx1GPBPY-Gqt8vFMsoTO0izDcj9V--86lP-9Ud-XFO9_OXVq5Ah5y36Fy9twbf9oetvxk_fn6CQVkJ_Nqrpe5KtJcyDayuLl9DTbVAKwaSzNU005xWj4HVufHbz7aF2IYZ17aKj_pkutwvKetejVQOVAK1GhVHcmR9tOJ5LI-YuDXFykNn6F05nnc_EaERp2QzOdYviATEiHiJtlrfdij5ktRf66Dq7lOz25XKLqirxof17PZFdDQpfhVaT6Gy60uggP6H513AVEGZW1DfIu02dR3t8eXhzYv8XMKm8bM93VQ9K_1gFnDyJtG02Sg%20token_type:Bearer%20id:101085123538940206503
	         ]})
	       ]%20
	     email:aokhotnikov@softjourn.com%20
	     name:Anatoliy%20Okhotnikov%20
	     given_name:Anatoliy%20
	     picture:https://lh5.googleusercontent.com/-fgrIHZSYhlI/AAAAAAAAAAI/AAAAAAAAAiI/ymKzDKpG2B4/photo.jpg%20
	     id:101085123538940206503%20
	     verified_email:%!s(bool=true)%20
	     link:https://plus.google.com/101085123538940206503]}#
	*/
	beego.Info(fmt.Sprintf("Email: %#v", user.Email()))

	//data := fmt.Sprintf("%#v", user)
	//io.WriteString(w, data)
	this.SetSession("email", user.Email())
	this.SetSession("name", user.Name())

	// redirect
	this.Redirect(fmt.Sprintf("/dashboard?user=%s", user), http.StatusFound)
}
