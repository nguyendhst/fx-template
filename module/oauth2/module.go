package oauth2

import (
	"github.com/nguyendhst/lagile/module/config"
	"go.uber.org/fx"
)

func GetModule(cfg *config.Config) fx.Option {
	var opts []fx.Option

	if cfg.Env.Secret.OAuth2.Provider.Google.ClientID != "" {
		opts = append(opts, fx.Provide(NewGoogleProvider))
	}

	return fx.Options(opts...)
}
