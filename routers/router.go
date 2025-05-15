package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	v1 "go_rest_api_template_with_gin/interfaces/handler/v1"
	"go_rest_api_template_with_gin/interfaces/middleware"
	"go_rest_api_template_with_gin/packages/env"
)

var Engin *gin.Engine

type Engine struct {
	Engine *gin.Engine
}

// SetBase initialize gin
func (e *Engine) SetBase() {
	// catch unexpected panic & return 500 error
	e.Engine.Use(gin.Recovery())

	// ヘルスチェック
	e.Engine.GET("/v1/status", HealthCheck)
}

// SetRouter setting routing infomation
func (e *Engine) SetRouter(v1 v1.AppHandler) {
	apiv1 := e.Engine.Group("/v1")

	apiv1.GET("/hoge",
		middleware.BindGetHogeRequestHeader())
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

// HealthCheck health check
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
