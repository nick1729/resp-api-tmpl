package config

import "time"

type Server struct {
	Host            string        `envconfig:"host" required:"true"`
	Port            string        `envconfig:"port" required:"true"`
	ReadTimeout     time.Duration `envconfig:"read_timeout" default:"30s"`
	WriteTimeout    time.Duration `envconfig:"write_timeout" default:"30s"`
	ShutdownTimeout time.Duration `envconfig:"shutdown_timeout" default:"30s"`
	Mode            string        `envconfig:"mode" required:"true"`
}
