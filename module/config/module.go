package config

import (
	"github.com/nguyendhst/fx-template/shared/constant"
	"go.uber.org/fx"
)

var Module = fx.Module(
	constant.CONFIG_MODULE,
	fx.Provide(New),
)
