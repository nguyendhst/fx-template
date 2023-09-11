package httpserver

import (
	"github.com/nguyendhst/lagile/shared/constant"

	"go.uber.org/fx"
)

var Module = fx.Module(constant.HTTPSERVER_MODULE, fx.Provide(
	NewEchoServer,
))
