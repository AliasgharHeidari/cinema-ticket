package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Server  ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type DatabaseConfig struct {
	DSN string `json:"dsn"`
}


func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var Cfg Config

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		panic(err)
	}

	return &Cfg, nil
}


func (s ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)

}
