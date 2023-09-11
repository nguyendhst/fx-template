package route

import (
	loginController "github.com/nguyendhst/clean-architecture-skeleton/api/controller/login"
	"github.com/nguyendhst/clean-architecture-skeleton/module/config"
	"github.com/nguyendhst/clean-architecture-skeleton/module/httpserver"
	"github.com/nguyendhst/clean-architecture-skeleton/shared/constant"
)

func NewAuthenticationRouter(env *config.Env, server *httpserver.Server, controller *loginController.LoginController) {
	authRoute := server.Prefix.Group(constant.AUTHENTICATION_ROUTE_PREFIX)

	// Login
	authRoute.POST("/login", controller.Login)
}
