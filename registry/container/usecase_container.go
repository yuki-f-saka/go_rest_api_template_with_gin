package container

import (
	"go_rest_api_template_with_gin/usecase"
)

// GetUserUsecase はユーザーユースケースのインスタンスを返します。
func GetUserUsecase() (usecase.UserUsecase, error) {
	userRepo, err := GetUserRepository()
	if err != nil {
		return nil, err
	}
	return usecase.NewUserUsecase(userRepo), nil
}
