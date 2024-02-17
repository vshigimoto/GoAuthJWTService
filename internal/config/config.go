package config

import "time"

type Config struct {
	Http Http
}

type Http struct {
	Port         int           `yaml:"port"`
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}
