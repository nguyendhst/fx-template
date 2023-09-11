package route

import (
	register "github.com/nguyendhst/clean-architecture-skeleton/api/controller/register"
	"github.com/nguyendhst/clean-architecture-skeleton/module/config"
	"github.com/nguyendhst/clean-architecture-skeleton/module/httpserver"
	"github.com/nguyendhst/clean-architecture-skeleton/shared/constant"
)

func NewPublicUserRouter(env *config.Env, server *httpserver.Server, controller *register.RegisterController) {
	userRoute := server.Prefix.Group(constant.USER_ROUTE_PREFIX)

	userRoute.POST("/register", controller.RegisterUser)

}
