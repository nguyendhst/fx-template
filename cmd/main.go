package main

import (
	"flag"

	"github.com/nguyendhst/lagile/api/route"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/core"
	"github.com/nguyendhst/lagile/module/database"
	"github.com/nguyendhst/lagile/module/httpserver"
	"github.com/nguyendhst/lagile/module/logger"
	"go.uber.org/fx"
)

func main() {
	env := flag.String("env", "local", "Environment")
	flag.Parse()

	cfg, err := config.New(*env)
	if err != nil {
		panic(err)
	}

	app := fx.New(
		core.GetModule(cfg.Env),
		logger.Module,
		database.GetModule(cfg.Env),
		route.Module,
		// fx.Invoke(route.NewRouter),

		fx.Invoke(func(server *httpserver.Server) error {
			go server.Start()
			return nil
		}),
	)
	app.Run()
}
