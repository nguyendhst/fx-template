package route

import (
	adminLogin "github.com/nguyendhst/lagile/api/controller/admin-login"
	register "github.com/nguyendhst/lagile/api/controller/register"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/httpserver"
	"go.uber.org/fx"
)

// var Module = fx.Provide(New)
// Setting up the router -- Does not need to return anything
var Module = fx.Invoke(New)

type (
	Params struct {
		fx.In

		Configs            *config.Config
		Server             *httpserver.Server
		LoginController    *adminLogin.AdminLoginController
		RegisterController *register.RegisterController
	}
)

func New(p Params) {
	// TODO: move this to config
	s := p.Server.SetPrefix("/api/v1")

	// Public API for authentication
	NewAuthenticationRouter(p.Configs, s, p.LoginController)
	NewPublicUserRouter(p.Configs, s, p.RegisterController)

	// Private API for user
	NewPrivateUserRouter(p.Configs, s)
}
