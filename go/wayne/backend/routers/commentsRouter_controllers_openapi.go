package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	const OpenAPIController = "k8s-lx1036/wayne/backend/controllers/openapi:OpenAPIController"
	beego.GlobalControllerRouter[OpenAPIController] = append(
		beego.GlobalControllerRouter[OpenAPIController],
		beego.ControllerComments{
			Method: "UpgradeDeployment",
			Router: `/upgrade_deployment`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Filters: nil,
			Params: nil,
		},
	)
}
