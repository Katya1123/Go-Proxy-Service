package config

import "github.com/gosidekick/goconfig"

type Config struct {
	Debug  bool   `cfgDefault:"true"`
	Server Server `cfg:"APP"`
}

type Server struct {
	Addr string `cfgDefault:"0.0.0.0"`
	Port int    `cfgDefault:"80"`
}

func InitConfig() (*Config, error) {
	config := Config{}

	err := goconfig.Parse(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
