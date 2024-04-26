package config

import (
	"fmt"
)

type Postgres struct {
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Host     string `envconfig:"HOST" required:"true"`
	Port     string `envconfig:"PORT" required:"true"`
	Database string `envconfig:"DATABASE" required:"true"`

	WithDebug bool `envconfig:"WITH_DEBUG" default:"false" `
}

const connStringTmpl = "postgres://%s:%s@%s:%s/%s"

func (p Postgres) ConnString() string {
	resp := fmt.Sprintf(
		connStringTmpl,
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Database,
	)

	return resp
}
