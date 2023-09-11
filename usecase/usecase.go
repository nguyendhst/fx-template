package usecase

import (
	"go.uber.org/fx"
)

// Provide your usecase implementations here.
func New() fx.Option {
	return fx.Provide(
		NewLoginUsecase,
		NewRegisterUsecase,
	)
}
