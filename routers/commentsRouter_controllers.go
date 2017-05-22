package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app-service/login-service/controllers:LoginController"] = append(beego.GlobalControllerRouter["app-service/login-service/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
