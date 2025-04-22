package registry

import (
	"go_rest_api_template_with_gin/routers"

	"github.com/gin-gonic/gin"
)

type Registry interface {
	Register() routers.Engine
}

type registry struct {
	engin routers.Engine
	// container container.Container
}

func NewRegistry() Registry {
	return &registry{
		engin: routers.Engine{
			Engine: gin.New(),
		},
	}
}

func (r registry) Register() routers.Engine {
	return r.engin
}
