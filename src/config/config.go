package config

import "github.com/kelseyhightower/envconfig"

var MainConfig *Config

type Config struct {
	DBConfig
}

// DB config struct
type DBConfig struct {
	DBHost   string `envconfig:"db_host" required:"true"`
	DBPort   string `envconfig:"db_port" required:"true"`
	DBName   string `envconfig:"db_name" required:"true"`
	DBSchema string `envconfig:"db_schema" required:"true"`
	DBUser   string `envconfig:"db_user" required:"true"`
	DBPass   string `envconfig:"db_pass" required:"true"`
	MaxConns int    `envconfig:"max_conns" required:"true"`
}

func InitConfigs() error {
	var cfg Config

	err := envconfig.Process("APP_CONFIG", &cfg)
	if err != nil {
		return err
	}

	return nil
}
