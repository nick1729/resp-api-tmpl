package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/errors"
)

type Config struct {
	Log      Log      `envconfig:"log" required:"true"`
	Server   Server   `envconfig:"server" required:"true"`
	Postgres Postgres `envconfig:"postgres" required:"true"`
	Redis    Redis    `envconfig:"redis" required:"true"`
	Version  string   `envconfig:"version" required:"true"`
}

const coreEnvPrefix = "my_app"

// Init - initialises config from .env file.
func Init() (Config, error) {
	var (
		cfg Config
		err error
	)

	err = godotenv.Load()
	if err != nil {
		return cfg, errors.Wrap(err, "loading .env")
	}

	err = envconfig.Process(coreEnvPrefix, &cfg)
	if err != nil {
		return cfg, errors.Wrap(err, "parsing config from env vars")
	}

	return cfg, nil
}
