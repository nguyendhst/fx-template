package auth

import (
	"context"

	u "github.com/nguyendhst/fx-template/domain/user"
)

type OauthLoginRequest struct {
	Email string `form:"email" binding:"required,email"`
}

type OauthLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type OauthLoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (u.User, error)
	CreateAccessToken(user *u.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *u.User, secret string, expiry int) (refreshToken string, err error)
}
