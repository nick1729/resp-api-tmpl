package config

import "fmt"

type Redis struct {
	Host string `envconfig:"host" required:"true"`
	Port string `envconfig:"port" required:"true"`
	Pass string `envconfig:"pass"`
	DB   int    `default:"0"      envconfig:"db"`
}

func (v *Redis) BuildAddressString() string {
	return fmt.Sprintf("%s:%s", v.Host, v.Port)
}
