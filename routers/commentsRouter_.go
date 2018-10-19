package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["test/controllers:BooklistController"] = append(beego.GlobalControllerRouter["test/controllers:BooklistController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams: param.Make(),
			Params: nil})

}
