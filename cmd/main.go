package main

import (
	"flag"

	"github.com/nguyendhst/fx-template/api/route"
	"github.com/nguyendhst/fx-template/module/config"
	"github.com/nguyendhst/fx-template/module/core"
	"github.com/nguyendhst/fx-template/module/database"
	"github.com/nguyendhst/fx-template/module/httpserver"
	"github.com/nguyendhst/fx-template/module/logger"
	"go.uber.org/fx"
)

func main() {
	env := flag.String("env", "development", "Environment")
	flag.Parse()

	if *env == "" {
		panic("Environment is required")
	}

	configs, err := config.New(*env)
	if err != nil {
		panic(err)
	}

	app := fx.New(
		core.GetModule(configs),
		logger.Module,
		database.GetModule(configs),
		route.Module,
		// fx.Invoke(route.NewRouter),

		fx.Invoke(func(server *httpserver.Server) error {
			go server.Start()
			return nil
		}),
	)
	app.Run()
}
