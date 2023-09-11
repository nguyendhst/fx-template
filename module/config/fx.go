package config

import (
	"github.com/nguyendhst/clean-architecture-skeleton/shared/constant"
	"go.uber.org/fx"
)

var Module = fx.Module(
	constant.CONFIG_MODULE,
	fx.Provide(NewEnv),
)
