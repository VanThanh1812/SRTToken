package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["srttoken/controllers:FaucetController"] = append(beego.GlobalControllerRouter["srttoken/controllers:FaucetController"],
		beego.ControllerComments{
			Method: "Balance",
			Router: `/balance`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srttoken/controllers:FaucetController"] = append(beego.GlobalControllerRouter["srttoken/controllers:FaucetController"],
		beego.ControllerComments{
			Method: "Buy",
			Router: `/buy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srttoken/controllers:UserController"] = append(beego.GlobalControllerRouter["srttoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srttoken/controllers:UserController"] = append(beego.GlobalControllerRouter["srttoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srttoken/controllers:UserController"] = append(beego.GlobalControllerRouter["srttoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srttoken/controllers:UserController"] = append(beego.GlobalControllerRouter["srttoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srttoken/controllers:UserController"] = append(beego.GlobalControllerRouter["srttoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
