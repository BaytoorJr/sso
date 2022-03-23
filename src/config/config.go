package config

import "github.com/kelseyhightower/envconfig"

var MainConfig *Config

type Config struct {
	DBConfig
}

// DB config struct
type DBConfig struct {
	PostgresMaxConns int    `envconfig:"postgres_max_conns" required:"true"`
	PostgresHost     string `envconfig:"postgres_host" required:"true"`
	PostgresPort     string `envconfig:"postgres_port" required:"true"`
	PostgresName     string `envconfig:"postgres_name" required:"true"`
	PostgresSchema   string `envconfig:"postgres_schema" required:"true"`
	PostgresUser     string `envconfig:"postgres_user" required:"true"`
	PostgresPass     string `envconfig:"postgres_pass" required:"true"`
}

func InitConfigs() error {
	var cfg Config

	err := envconfig.Process("APP_CONFIG", &cfg)
	if err != nil {
		return err
	}

	MainConfig = &cfg

	return nil
}
