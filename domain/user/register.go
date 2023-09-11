package user

import (
	"context"

	"github.com/nguyendhst/lagile/domain/response"
)

type (
	UserRegisterRequest struct {
		Name     string `json:"name" form:"name" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserRegisterResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

type RegisterUsecase interface {
	NewResponse(user *UserRegisterResponse) response.Response
	RegisterUser(c context.Context, user *User) (*User, error)
}
