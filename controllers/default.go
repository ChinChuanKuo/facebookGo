package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.tpl"
	if this.Ctx.Request.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, this.Ctx.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		return
	}
	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create request: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response: %s\n", err.Error())
		return
	}
	fmt.Fprintf(this.Ctx.ResponseWriter, "Response: %s", content)
}
