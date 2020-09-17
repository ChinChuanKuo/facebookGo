package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
)

type CallbackController struct {
	beego.Controller
}

func (this *CallbackController) Get() {
	if this.Ctx.Request.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, this.Ctx.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create request: %s\n", err.Error())
		http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response: %s\n", err.Error())
		http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(this.Ctx.ResponseWriter, "Response: %s", content)
}
