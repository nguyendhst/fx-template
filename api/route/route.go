package route

import (
	login "github.com/nguyendhst/clean-architecture-skeleton/api/controller/login"
	register "github.com/nguyendhst/clean-architecture-skeleton/api/controller/register"
	"github.com/nguyendhst/clean-architecture-skeleton/module/config"
	"github.com/nguyendhst/clean-architecture-skeleton/module/httpserver"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRouter)

type (
	Params struct {
		fx.In

		Env                *config.Env
		Server             *httpserver.Server
		LoginController    *login.LoginController
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
