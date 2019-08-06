package spi

import (
	"github.com/BurntSushi/toml"
	"log"
)

type AppConfig struct {
	Database DatabaseConfig
	Http     HttpConfig
}

type DatabaseConfig struct {
	Driver string
	Url    string
	Host   string
	Port   int
}

type HttpConfig struct {
	Host string
	Port int
}

func Load(file string) (*AppConfig, error) {
	var config = AppConfig{}

	if _, err := toml.DecodeFile(file, &config); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &config, nil
}
