package core

import (
	"github.com/nguyendhst/fx-template/api/controller"
	"github.com/nguyendhst/fx-template/module/config"
	"github.com/nguyendhst/fx-template/module/httpserver"
	"github.com/nguyendhst/fx-template/repository"
	"github.com/nguyendhst/fx-template/usecase"
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
