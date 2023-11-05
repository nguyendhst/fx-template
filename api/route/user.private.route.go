package route

import (
	"github.com/nguyendhst/fx-template/api/middleware"
	"github.com/nguyendhst/fx-template/module/config"
	"github.com/nguyendhst/fx-template/module/httpserver"
	"github.com/nguyendhst/fx-template/shared/constant"
)

func NewPrivateUserRouter(cfg *config.Config, server *httpserver.Server) {
	userRoute := server.Prefix.Group(constant.USER_ROUTE_PREFIX)

	// Apply middleware
	userRoute.Use(middleware.JWTMiddleware(cfg.Env.Secret.JwtSecret.Access.Key))
	// Login
}
