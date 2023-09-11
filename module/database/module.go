package database

import (
	"github.com/nguyendhst/lagile/shared/constant"
	"go.uber.org/fx"
)

var Module = fx.Module(
	constant.DATABASE_MODULE,
	fx.Provide(
		fx.Annotate(
			NewPostgresDatabase,
			fx.As(new(Database)),
		),
	),
)
