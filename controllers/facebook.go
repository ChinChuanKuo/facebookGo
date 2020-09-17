package controllers

import (
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type GoogleController struct {
	beego.Controller
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		Scopes:       []string{"public_profile,email"},
		Endpoint:     facebook.Endpoint,
	}
	//TODO randomize it
	randomState = "random"
)

func (this *GoogleController) Get() {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, url, http.StatusTemporaryRedirect)
}
