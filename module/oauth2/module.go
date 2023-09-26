package oauth2

import (
	"github.com/nguyendhst/lagile/module/config"
	"go.uber.org/fx"
)

func GetModule(env *config.Env) fx.Option {
	var opts []fx.Option

	if env.Secret.OAuth2.Provider.Google.ClientID != "" {
		opts = append(opts, fx.Provide(NewGoogleProvider))
	}

	return fx.Options(opts...)
}
