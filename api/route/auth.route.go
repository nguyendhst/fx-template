package route

import (
	loginController "github.com/nguyendhst/fx-template/api/controller/admin-login"
	"github.com/nguyendhst/fx-template/module/config"
	"github.com/nguyendhst/fx-template/module/httpserver"
	"github.com/nguyendhst/fx-template/shared/constant"
)

func NewAuthenticationRouter(cfg *config.Config, server *httpserver.Server, controller *loginController.AdminLoginController) {
	authRoute := server.Prefix.Group(constant.AUTHENTICATION_ROUTE_PREFIX)

	// Login
	authRoute.POST("/login", controller.Login)
}
