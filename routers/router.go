package routers

import (
	"go_rest_api_template_with_gin/packages/env"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Engin *gin.Engine

type Engine struct {
	Engine *gin.Engine
}

// SetBase initialize gin
func (e *Engine) SetBase() {
	// catch unexpected panic & return 500 error
	e.Engine.Use(gin.Recovery())
}

// SetCORS cors infomation
func (e *Engine) SetCORS() {
	e.Engine.Use(cors.New(cors.Config{
		// AllowAllOrigins: false,
		AllowOrigins: []string{
			env.Env().AccessAllowOrigin,
			env.Env().AccessAllowOriginWeb,
			env.Env().AccessAllowOriginCms,
		},
		// AllowOriginFunc: func(origin string) bool {
		// 	panic("TODO")
		// },
		// AllowOriginWithContextFunc: func(c *gin.Context, origin string) bool {
		// 	panic("TODO")
		// },
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
			"OPTIONS",
		},
		// AllowPrivateNetwork:       false,
		AllowHeaders: []string{
			"*",
		},
		// AllowCredentials:          false,
		// ExposeHeaders:             []string{},
		MaxAge: 24 * time.Hour,
		// AllowWildcard:             false,
		// AllowBrowserExtensions:    false,
		// CustomSchemas:             []string{},
		// AllowWebSockets:           false,
		// AllowFiles:                false,
		// OptionsResponseStatusCode: 0,
	}))
}
