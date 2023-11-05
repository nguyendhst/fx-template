package core

import (
	"github.com/nguyendhst/lagile/api/controller"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/repository"
	"github.com/nguyendhst/lagile/usecase"
	"go.uber.org/fx"
)

func GetModule(cfg *config.Config) fx.Option {
	return fx.Options(
		fx.Supply(cfg),
		httpserver.Module,
		repository.Module,
		usecase.Module,
		controller.New(),
	)
}
