package main

import (
	_ "github.com/lib/pq"
	"github.com/nguyendhst/lagile/api/controller"
	"github.com/nguyendhst/lagile/api/route"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/database"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/module/logger"
	"github.com/nguyendhst/lagile/repository"
	"github.com/nguyendhst/lagile/usecase"
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
