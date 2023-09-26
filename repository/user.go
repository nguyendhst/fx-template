package repository

import (
	"context"

	domain "github.com/nguyendhst/lagile/domain/user"
)

type (
	UserRepository interface {
		Create(ctx context.Context, user *domain.User) (createdUser *domain.User, err error)
		Fetch(ctx context.Context) (users []domain.User, err error)
		GetByEmail(ctx context.Context, email string) (domain.User, error)
		GetByID(ctx context.Context, id int32) (domain.User, error)
	}
)
