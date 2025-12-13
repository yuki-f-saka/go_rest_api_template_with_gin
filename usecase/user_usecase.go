package usecase

import (
	"context"
	"go_rest_api_template_with_gin/domain"
	"go_rest_api_template_with_gin/infrastructure/repository"
)

// UserUsecase はユーザーに関するユースケースのインターフェースです。
type UserUsecase interface {
	CreateUser(ctx context.Context, user *domain.User) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase は UserUsecase の新しいインスタンスを返します。
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

// CreateUser は新しいユーザーを作成します。
func (u *userUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	return u.userRepo.Save(ctx, user)
}
