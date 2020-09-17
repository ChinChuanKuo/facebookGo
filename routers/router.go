package routers

import (
	"facebookGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/facebook", &controllers.GoogleController{})
	beego.Router("/callback", &controllers.CallbackController{})
}
