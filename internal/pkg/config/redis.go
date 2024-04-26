package config

type Redis struct {
	Host string `envconfig:"host" required:"true"`
	Port string `envconfig:"port" required:"true"`
	Pass string `envconfig:"pass"`
	DB   int    `envconfig:"db" default:"0"`
}
