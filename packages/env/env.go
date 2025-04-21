package env

import (
	"github.com/kelseyhightower/envconfig"
)

var cfg Config

type Config struct {
	GoEnv      string `envconfig:"GO_ENV" default:"prd"`
	ListenHost string `envconfig:"LISTEN_HOST"`
	ListenPort string `envconfig:"LISTEN_PORT"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

func setupEnv() error {
	return envconfig.Process("", &cfg)
}

func Env() Config {
	return cfg
}
