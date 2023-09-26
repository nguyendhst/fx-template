package database

import (
	"github.com/nguyendhst/lagile/module/config"
	_ "github.com/nguyendhst/lagile/shared/constant"
	"go.uber.org/fx"
)

func GetModule(env config.Env) fx.Option {
	var opts []fx.Option

	if env.Database.Postgres.Host != "" {
		opts = append(opts, fx.Provide(NewPostgresClient))
	}
	if env.Database.Mongo.Host != "" {
		opts = append(opts, fx.Provide(NewMongoClient))
	}

	return fx.Options(opts...)
}
