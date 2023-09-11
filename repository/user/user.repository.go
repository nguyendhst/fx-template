package user

import (
	"context"
	"database/sql"

	domain "github.com/nguyendhst/clean-architecture-skeleton/domain/user"
	db "github.com/nguyendhst/clean-architecture-skeleton/module/database"
	"github.com/nguyendhst/clean-architecture-skeleton/shared/util"
	sqlc "github.com/nguyendhst/clean-architecture-skeleton/sqlc/generated"
)

type userRepository struct {
	table   string
	db      *sql.DB
	queries *sqlc.Queries
}

func New(db db.Database) domain.UserRepository {
	return &userRepository{
		table:   "users",
		db:      db.GetDBConnection(),
		queries: db.GetQueries(),
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) (createdUser *domain.User, err error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	res, err := ur.queries.CreateUser(ctx, sqlc.CreateUserParams{
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
	query := "SELECT * FROM users"
	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	query := "SELECT * FROM users WHERE email = $1"
	err = ur.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id string) (user domain.User, err error) {
	query := "SELECT * FROM users WHERE id = $1"
	err = ur.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
