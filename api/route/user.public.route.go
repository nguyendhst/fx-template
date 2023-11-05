package route

import (
	register "github.com/nguyendhst/fx-template/api/controller/register"
	"github.com/nguyendhst/fx-template/module/config"
	"github.com/nguyendhst/fx-template/module/httpserver"
	"github.com/nguyendhst/fx-template/shared/constant"
)

func NewPublicUserRouter(cfg *config.Config, server *httpserver.Server, controller *register.RegisterController) {
	userRoute := server.Prefix.Group(constant.USER_ROUTE_PREFIX)

	userRoute.POST("/register", controller.RegisterUser)
}
