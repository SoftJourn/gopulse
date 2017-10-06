package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
//	"io"
	"net/http"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

type LoginController struct {
	beego.Controller
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
		beego.Error(err.Error())
		this.Redirect("/", http.StatusInternalServerError)
		return
	}

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
