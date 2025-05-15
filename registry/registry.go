package registry

import (
	"go_rest_api_template_with_gin/registry/container"
	"go_rest_api_template_with_gin/routers"

	"github.com/gin-gonic/gin"
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
	r.engin.SetRouter(r.container.GetAppHandler)
	return r.engin
}
