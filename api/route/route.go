package route

import (
	adminLogin "github.com/nguyendhst/lagile/api/controller/admin-login"
	register "github.com/nguyendhst/lagile/api/controller/register"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/httpserver"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRouter)

type (
	Params struct {
		fx.In

		Env                *config.Env
		Server             *httpserver.Server
		LoginController    *adminLogin.AdminLoginController
		RegisterController *register.RegisterController
	}
)

func NewRouter(p Params) {
	s := p.Server.SetPrefix("/api/v1")

	// Public API for authentication
	NewAuthenticationRouter(p.Env, s, p.LoginController)
	NewPublicUserRouter(p.Env, s, p.RegisterController)

	// Private API for user
	NewPrivateUserRouter(p.Env, s)
}
