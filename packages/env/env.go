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

	AccessAllowOrigin    string `envconfig:"ACCESS_ALLOW_ORIGIN"`
	AccessAllowOriginWeb string `envconfig:"ACCESS_ALLOW_ORIGIN_WEB"`
	AccessAllowOriginCms string `envconfig:"ACCESS_ALLOW_ORIGIN_CMS"`
}

func SetupEnv() error {
	return envconfig.Process("", &cfg)
}

func Env() Config {
	return cfg
}
