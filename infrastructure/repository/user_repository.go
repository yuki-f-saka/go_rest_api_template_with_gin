package repository

import (
	"context"
	"database/sql"
	"go_rest_api_template_with_gin/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(
	ctx context.Context,
	user *domain.User,
) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO users (id, name, email) VALUES (?, ?, ?)`,
		user.ID,
		user.Name,
		user.Email,
	)
	return err
}
