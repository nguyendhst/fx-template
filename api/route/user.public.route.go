package route

import (
	register "github.com/nguyendhst/lagile/api/controller/register"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/shared/constant"
)

func NewPublicUserRouter(env *config.Env, server *httpserver.Server, controller *register.RegisterController) {
	userRoute := server.Prefix.Group(constant.USER_ROUTE_PREFIX)

	userRoute.POST("/register", controller.RegisterUser)
}
