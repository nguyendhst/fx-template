package main

import (
	_ "github.com/lib/pq"
	"github.com/nguyendhst/clean-architecture-skeleton/api/controller"
	"github.com/nguyendhst/clean-architecture-skeleton/api/route"
	"github.com/nguyendhst/clean-architecture-skeleton/module/config"
	"github.com/nguyendhst/clean-architecture-skeleton/module/database"
	"github.com/nguyendhst/clean-architecture-skeleton/module/httpserver"
	"github.com/nguyendhst/clean-architecture-skeleton/module/logger"
	"github.com/nguyendhst/clean-architecture-skeleton/repository"
	"github.com/nguyendhst/clean-architecture-skeleton/usecase"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		logger.Module,
		config.Module,
		database.Module,
		httpserver.Module,
		usecase.New(),
		repository.New(),
		controller.New(),
		fx.Invoke(route.NewRouter),

		fx.Invoke(func(server *httpserver.Server) error {
			go server.Start()
			return nil
		}),
	)
	app.Run()

}
