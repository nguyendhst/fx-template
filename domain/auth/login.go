package auth

import (
	"context"

	u "github.com/nguyendhst/clean-architecture-skeleton/domain/user"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (u.User, error)
	CreateAccessToken(user *u.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *u.User, secret string, expiry int) (refreshToken string, err error)
}
