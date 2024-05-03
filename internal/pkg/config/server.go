package config

import "time"

type Server struct {
	Host            string        `envconfig:"host"       required:"true"`
	Port            string        `envconfig:"port"       required:"true"`
	BodyLimit       int           `envconfig:"body_limit" required:"true"`
	ReadTimeout     time.Duration `default:"30s"          envconfig:"read_timeout"`
	WriteTimeout    time.Duration `default:"30s"          envconfig:"write_timeout"`
	ShutdownTimeout time.Duration `default:"30s"          envconfig:"shutdown_timeout"`
	Mode            string        `envconfig:"mode"       required:"true"`
}
