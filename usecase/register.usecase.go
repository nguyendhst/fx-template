package usecase

import (
	"context"
	"time"

	response "github.com/nguyendhst/clean-architecture-skeleton/domain/response"
	domain "github.com/nguyendhst/clean-architecture-skeleton/domain/user"
)

type (
	registerUsecase struct {
		userRepository domain.UserRepository
		contextTimeout time.Duration
	}
)

func NewRegisterUsecase(userRepository domain.UserRepository) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		contextTimeout: time.Millisecond * 5000,
	}
}

func (ru *registerUsecase) NewResponse(user *domain.UserRegisterResponse) response.Response {
	meta := response.NewResponseMeta(200, "OK", "")
	data := &domain.UserRegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return response.NewResponse(meta, data)
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
