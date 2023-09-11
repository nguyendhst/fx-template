package main

import (
	_ "github.com/lib/pq"
	"github.com/nguyendhst/lagile/api/route"
	"github.com/nguyendhst/lagile/internal/core"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/database"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/module/logger"
	"go.uber.org/fx"
)

func main() {

	cfg, err := config.NewEnv()
	if err != nil {
		panic(err)
	}

	app := fx.New(
		core.GetModule(*cfg),
		logger.Module,
		config.Module,
		database.Module,
		fx.Invoke(route.NewRouter),

		fx.Invoke(func(server *httpserver.Server) error {
			go server.Start()
			return nil
		}),
	)
	app.Run()

}
