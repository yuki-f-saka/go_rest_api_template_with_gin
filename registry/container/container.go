package container

import (
	v1 "go_rest_api_template_with_gin/interfaces/handler/v1"
)

// Container DIコンテナの枠となる構造体
type Container struct{}

// NewAppHandler はアプリケーションハンドラーの新しいインスタンスを返します。
func (c Container) NewAppHandler() (*v1.AppHandler, error) {
	return GetAppHandler()
}
