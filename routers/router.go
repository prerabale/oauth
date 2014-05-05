package routers

import (
	"github.com/arkors/oauth/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/app", &controllers.AppController{}, "post:CreateApp")
	beego.Router("/v1/app/:id", &controllers.AppController{}, "get:GetAppKey")
	beego.Router("/v1/app/:id", &controllers.AppController{}, "put:ResetAppKey")
	beego.Router("/v1/token", &controllers.TokenController{}, "post:CreateToken")
	beego.Router("/v1/token/:token", &controllers.TokenController{}, "get:VerifyToken")
	beego.Router("/", &controllers.MainController{})
}
