package user

import (
	"context"
	"strconv"

	domain "github.com/nguyendhst/lagile/domain/user"
	"github.com/nguyendhst/lagile/shared/util"
	sqlc "github.com/nguyendhst/lagile/sqlc/generated"
)

type userRepository struct {
	querier *sqlc.Queries
}

func NewRepository(db *sqlc.Queries) *userRepository {
	return &userRepository{
		db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) (createdUser *domain.User, err error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// TODO: Convert to TX
	res, err := ur.querier.CreateUser(ctx, sqlc.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    string(res.ID),
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func (ur *userRepository) Fetch(ctx context.Context) (users []domain.User, err error) {
	res, err := ur.querier.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range res {
		users = append(users, domain.User{
			ID:    string(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return users, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	res, err := ur.querier.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:    strconv.Itoa(int(res.ID)),
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id int32) (user domain.User, err error) {
	res, err := ur.querier.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:    strconv.Itoa(int(res.ID)),
		Name:  res.Name,
		Email: res.Email,
	}, nil
}
