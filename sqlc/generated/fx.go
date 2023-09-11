package sqlc

import "go.uber.org/fx"

var Module = fx.Module(
	"sqlc-module",
	fx.Provide(
		New,
	),
)
