package route

import (
	"github.com/nguyendhst/lagile/api/middleware"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/shared/constant"
)

func NewPrivateUserRouter(cfg *config.Config, server *httpserver.Server) {
	userRoute := server.Prefix.Group(constant.USER_ROUTE_PREFIX)

	// Apply middleware
	userRoute.Use(middleware.JWTMiddleware(cfg.Env.Secret.JwtSecret.Access.Key))
	// Login
}
