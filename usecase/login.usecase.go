package usecase

import (
	"context"
	"time"

	auth "github.com/nguyendhst/lagile/domain/auth"
	user "github.com/nguyendhst/lagile/domain/user"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/repository"
	"github.com/nguyendhst/lagile/shared/util"
)

type loginUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewAdminLoginUsecase(userRepository repository.UserRepository, env *config.Env) auth.BasicLoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: time.Millisecond * time.Duration(env.App.Login.Timeout),
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (user.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *user.User, secret string, expiry time.Duration) (accessToken string, err error) {
	return util.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *user.User, secret string, expiry time.Duration) (refreshToken string, err error) {
	return util.CreateRefreshToken(user, secret, expiry)
}
