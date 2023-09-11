package route

import (
	"github.com/nguyendhst/lagile/api/middleware"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/shared/constant"
)

func NewPrivateUserRouter(env *config.Env, server *httpserver.Server) {
	userRoute := server.Prefix.Group(constant.USER_ROUTE_PREFIX)

	// Apply middleware
	userRoute.Use(middleware.JWTMiddleware(env.AccessTokenSecret))
	// Login

}
