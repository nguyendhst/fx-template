package controller

import (
	loginController "github.com/nguyendhst/clean-architecture-skeleton/api/controller/login"
	registerController "github.com/nguyendhst/clean-architecture-skeleton/api/controller/register"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		// Provide your controller implementations here.
		loginController.New,
		registerController.New,
	)
}
