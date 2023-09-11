package route

import (
	"github.com/nguyendhst/clean-architecture-skeleton/api/middleware"
	"github.com/nguyendhst/clean-architecture-skeleton/module/config"
	"github.com/nguyendhst/clean-architecture-skeleton/module/httpserver"
	"github.com/nguyendhst/clean-architecture-skeleton/shared/constant"
)

func NewPrivateUserRouter(env *config.Env, server *httpserver.Server) {
	userRoute := server.Prefix.Group(constant.USER_ROUTE_PREFIX)

	// Apply middleware
	userRoute.Use(middleware.JWTMiddleware(env.AccessTokenSecret))
	// Login

}
