package registry

import (
	"go_rest_api_template_with_gin/registry/container"
	"go_rest_api_template_with_gin/routers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go_rest_api_template_with_gin/packages/log"
)

type Registry interface {
	Register() routers.Engine
}

type registry struct {
	engin     routers.Engine
	container container.Container
}

func NewRegistry() Registry {
	return &registry{
		engin: routers.Engine{
			// gin.Default() is not used because it has a logger and recovery middleware
			Engine: gin.New(),
		},
		container: container.Container{},
	}
}

func (r registry) Register() routers.Engine {
	r.engin.SetBase()
	r.engin.SetCORS()

	// ハンドラーの初期化
	appHandler, err := r.container.NewAppHandler()
	if err != nil {
		log.Logger.Fatal("failed to initialize app handler", zap.Error(err))
	}

	r.engin.SetRouter(*appHandler)
	return r.engin
}
