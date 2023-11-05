package controller

import (
	adminLoginController "github.com/nguyendhst/fx-template/api/controller/admin-login"
	registerController "github.com/nguyendhst/fx-template/api/controller/register"
	"go.uber.org/fx"
)

// Provide your controller implementations here.
func New() fx.Option {
	return fx.Provide(
		adminLoginController.New,
		registerController.New,
	)
}
