package container

import (
	"go_rest_api_template_with_gin/infrastructure/repository"
)

// GetUserRepository はユーザーリポジトリのインスタンスを返します。
func GetUserRepository() (repository.UserRepository, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	return repository.NewUserRepository(db), nil
}
