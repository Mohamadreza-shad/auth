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
	ResourcePath string `required:"true" envconfig:"GO_HEXAGONAL_I18N_RESOURCE_PATH"`
	Languages    string `required:"true" envconfig:"GO_HEXAGONAL_I18N_LANGUAGES"`
}

type Graylog struct {
	Format         string `required:"true" envconfig:"GO_HEXAGONAL_GRAYLOG_FORMAT"`
	Level          string `required:"true" envconfig:"GO_HEXAGONAL_GRAYLOG_LEVEL"`
	NoColor        bool   `required:"true" envconfig:"GO_HEXAGONAL_GRAYLOG_NOCOLOR"`
	GraylogAddress string `required:"true" envconfig:"GO_HEXAGONAL_GRAYLOG_ADDRESS"`
	ServiceName    string `required:"true" envconfig:"GO_HEXAGONAL_GRAYLOG_SERVICE_NAME"`
	Version        string `required:"true" envconfig:"GO_HEXAGONAL_GRAYLOG_VERSION"`
	Facility       string `required:"true" envconfig:"GO_HEXAGONAL_GRAYLOG_FACILITY"`
}

type Database struct {
	DSN          string `required:"true" envconfig:"GO_HEXAGONAL_DATABASE_DNS"`
	MaxOpenConns int    `required:"true" envconfig:"GO_HEXAGONAL_DATABASE_MAX_OPEN_CONNECTION"`
	MaxIdleConns int    `required:"true" envconfig:"GO_HEXAGONAL_DATABASE_MAX_IDLE_CONNECTION"`
}

type Server struct {
	HttpAddress   string `required:"true" envconfig:"GO_HEXAGONAL_SERVER_HTTP_ADDRESS"`
	GrpcAddress   string `required:"true" envconfig:"GO_HEXAGONAL_SERVER_GRPC_ADDRESS"`
	HealthAddress string `required:"true" envconfig:"GO_HEXAGONAL_SERVER_HEALTH_ADDRESS"`
}

func GetConfig() (Config, error) {
	var c Config
	err := envconfig.Process("GO_HEXAGONAL_", &c)
	return c, err
}
