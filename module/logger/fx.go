package logger

import (
	"github.com/nguyendhst/clean-architecture-skeleton/shared/constant"

	"go.uber.org/fx"
)

var Module = fx.Module(
	constant.LOGGER_MODULE,
	fx.Provide(
		fx.Annotate(
			NewLogger,
			fx.As(new(Logger)),
		),
	),
)
