package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["last/controllers:LibController"] = append(beego.GlobalControllerRouter["last/controllers:LibController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["last/controllers:LibController"] = append(beego.GlobalControllerRouter["last/controllers:LibController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["last/controllers:LibController"] = append(beego.GlobalControllerRouter["last/controllers:LibController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["last/controllers:LibController"] = append(beego.GlobalControllerRouter["last/controllers:LibController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["last/controllers:TargetController"] = append(beego.GlobalControllerRouter["last/controllers:TargetController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["last/controllers:TargetController"] = append(beego.GlobalControllerRouter["last/controllers:TargetController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
