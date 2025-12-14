package container

import (
	v1 "go_rest_api_template_with_gin/interfaces/handler/v1"
)

// GetAppHandler はアプリケーションハンドラーのインスタンスを返します。
func GetAppHandler() (*v1.AppHandler, error) {
	userUsecase, err := GetUserUsecase()
	if err != nil {
		return nil, err
	}
	return v1.NewAppHandler(userUsecase), nil
}
