package config

import (
	"github.com/nguyendhst/lagile/shared/constant"
	"go.uber.org/fx"
)

var Module = fx.Module(
	constant.CONFIG_MODULE,
	fx.Provide(New),
)
