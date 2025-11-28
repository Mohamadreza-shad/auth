package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env      string `required:"true" envconfig:"GO_HEXAGONAL_ENV"`
	Database Database
	Server   Server
	I18n     I18n
}

type I18n struct {
	ResourcePath string `required:"true" envconfig:"I18N_RESOURCE_PATH"`
	Languages    string `required:"true" envconfig:"I18N_LANGUAGES"`
}

type Database struct {
	DSN          string `required:"true" envconfig:"DATABASE_DNS"`
	MaxOpenConns int    `required:"true" envconfig:"DATABASE_MAX_OPEN_CONNECTION"`
	MaxIdleConns int    `required:"true" envconfig:"DATABASE_MAX_IDLE_CONNECTION"`
}

type Server struct {
	HttpAddress   string `required:"true" envconfig:"SERVER_HTTP_ADDRESS"`
	HealthAddress string `required:"true" envconfig:"SERVER_HEALTH_ADDRESS"`
}

func GetConfig() (Config, error) {
	var c Config
	err := envconfig.Process("GO_HEXAGONAL_", &c)
	return c, err
}
