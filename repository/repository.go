package repository

import (
	userRepository "github.com/nguyendhst/lagile/repository/user"
	"go.uber.org/fx"
)

// Provide your repository implementations here.
func New() fx.Option {
	return fx.Provide(
		userRepository.New,
	)
}
