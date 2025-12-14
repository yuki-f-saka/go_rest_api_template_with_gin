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

	DBDriver   string `envconfig:"DB_DRIVER" default:"mysql"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"3306"`
	DBUser     string `envconfig:"DB_USER" default:"root"`
	DBPassword string `envconfig:"DB_PASSWORD" default:""`
	DBName     string `envconfig:"DB_NAME" default:"app"`
}

func SetupEnv() error {
	return envconfig.Process("", &cfg)
}

func Env() Config {
	return cfg
}
