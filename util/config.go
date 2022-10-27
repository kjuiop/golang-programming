package util

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
}

func ConfInitialize() (*Config, error) {
	c := new(Config)

	err := envconfig.Process("reporter", c)
	if err != nil {
		log.Println("[ConfInitialize] failed read config :", err)
		return nil, err
	}

	return c, nil
}
