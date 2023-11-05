package route

import (
	loginController "github.com/nguyendhst/lagile/api/controller/admin-login"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/shared/constant"
)

func NewAuthenticationRouter(cfg *config.Config, server *httpserver.Server, controller *loginController.AdminLoginController) {
	authRoute := server.Prefix.Group(constant.AUTHENTICATION_ROUTE_PREFIX)

	// Login
	authRoute.POST("/login", controller.Login)
}
