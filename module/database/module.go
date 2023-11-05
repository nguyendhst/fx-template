package database

import (
	"github.com/nguyendhst/fx-template/module/config"
	_ "github.com/nguyendhst/fx-template/shared/constant"
	"go.uber.org/fx"
)

func GetModule(cfg *config.Config) fx.Option {
	var opts []fx.Option

	if cfg.Env.Database.Postgres.Host != "" {
		opts = append(opts, fx.Provide(NewPostgresClient))
	}
	if cfg.Env.Database.Mongo.Host != "" {
		opts = append(opts, fx.Provide(NewMongoClient))
	}

	return fx.Options(opts...)
}
