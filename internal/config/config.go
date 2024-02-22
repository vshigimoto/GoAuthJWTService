package config

import "time"

type Config struct {
	Http  Http
	Mongo Mongo
}

type Http struct {
	Port         int           `yaml:"port"`
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type Mongo struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Url        string `yaml:"url"`
	DbName     string `yaml:"db"`
	Collection string `yaml:"collection"`
}
