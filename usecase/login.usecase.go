package usecase

import (
	"context"
	"time"

	auth "github.com/nguyendhst/lagile/domain/auth"
	user "github.com/nguyendhst/lagile/domain/user"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/shared/util"
)

type loginUsecase struct {
	userRepository user.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository user.UserRepository, env *config.Env) auth.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: time.Millisecond * time.Duration(env.LoginUsecaseTimeout_MS),
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (user.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *user.User, secret string, expiry int) (accessToken string, err error) {
	return util.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *user.User, secret string, expiry int) (refreshToken string, err error) {
	return util.CreateRefreshToken(user, secret, expiry)
}
