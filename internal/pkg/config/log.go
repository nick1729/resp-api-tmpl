package config

type Log struct {
	Level string `envconfig:"level" required:"true"`
}
