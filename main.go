package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"go_rest_api_template_with_gin/packages/env"
	"go_rest_api_template_with_gin/packages/log"
)

var version string

func main() {
	setupEnv()
	setupLogger()
	addr := fmt.Sprintf("%s:%s",
		env.Env().ListenHost,
		env.Env().ListenPort,
	)

	// router
	router := setupRouter()

	// server
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// graceful restart & shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("failed to serve server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server forced to shutdown:", zap.Error(err))
	}

	log.Logger.Info("Server exiting")
}

func setupRouter() *gin.Engine {
	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API Routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	return router
}

func readEnv() {
	if err := godotenv.Load(fmt.Sprintf("./config/%s.env", os.Getenv("GO_ENV"))); err != nil {
		panic(err)
	}
}

func setupEnv() {
	readEnv()
}

func setupLogger() {
	lg, err := log.NewLogger(&log.Config{
		Env:        os.Getenv("GO_ENV"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		AppName:    os.Getenv("APP_NAME"),
		AppVersion: version,
	})
	if err != nil {
		panic(err)
	}
	log.SetLogger(lg)
}
