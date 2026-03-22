package main

import (
	"context"
	"fmt"
	"log/slog"
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
	"go_rest_api_template_with_gin/registry"
	"go_rest_api_template_with_gin/server"
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

	// graceful shutdown orchestrator
	// 合計タイムアウト 20s: http-shutdown 15s
	orch := server.NewShutdownOrchestrator(20 * time.Second)
	orch.Register("http-shutdown", 15*time.Second, func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	})

	// SIGINT・SIGTERM を受信するチャネルを登録する
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("failed to serve server", zap.Error(err))
		}
	}()

	sig := <-sigCh
	log.Logger.Info("gracefully stopping server...", zap.String("signal", sig.String()))

	// 2回目のシグナルで強制終了
	go func() {
		sig := <-sigCh
		log.Logger.Warn("forced exit", zap.String("signal", sig.String()))
		os.Exit(1)
	}()

	if err := orch.Shutdown(slog.Default()); err != nil {
		log.Logger.Error("shutdown error", zap.Error(err))
	}
	log.Logger.Info("stopped server")
}

func setupRouter() *gin.Engine {
	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// DI注入
	reg := registry.NewRegistry()
	engine := reg.Register()
	router := engine.Engine

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
