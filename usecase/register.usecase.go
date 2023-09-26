package usecase

import (
	"context"
	"time"

	domain "github.com/nguyendhst/lagile/domain/user"
	"github.com/nguyendhst/lagile/repository"
)

type (
	registerUsecase struct {
		userRepository repository.UserRepository
		contextTimeout time.Duration
	}
)

func NewRegisterUsecase(userRepository repository.UserRepository) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		contextTimeout: time.Millisecond * 5000,
	}
}

func (ru *registerUsecase) RegisterUser(c context.Context, u *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()

	user := &domain.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	return ru.userRepository.Create(ctx, user)
}
