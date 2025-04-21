package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"go_rest_api_template_with_gin/packages/log"
)

var version string

func main() {
	setupEnv()
	setupLogger()

	// registry
	// router
	// server

	// graceful restart & shutdown
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
