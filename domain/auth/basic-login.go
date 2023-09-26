package auth

import (
	"context"
	"time"

	u "github.com/nguyendhst/lagile/domain/user"
)

type BasicLoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type BasicLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type BasicLoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (u.User, error)
	CreateAccessToken(user *u.User, secret string, expiry time.Duration) (accessToken string, err error)
	CreateRefreshToken(user *u.User, secret string, expiry time.Duration) (refreshToken string, err error)
}
