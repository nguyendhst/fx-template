package repository

import (
	userRepository "github.com/nguyendhst/clean-architecture-skeleton/repository/user"
	"go.uber.org/fx"
)

// Provide your repository implementations here.
func New() fx.Option {
	return fx.Provide(
		userRepository.New,
	)
}
