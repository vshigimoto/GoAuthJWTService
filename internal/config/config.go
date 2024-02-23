package config

import "time"

type Config struct {
	HTTP  HTTP
	Mongo Mongo
}

type HTTP struct {
	Port            int           `yaml:"port"`
	Address         string        `yaml:"address"`
	ReadTimeout     time.Duration `yaml:"readTimeout"`
	WriteTimeout    time.Duration `yaml:"writeTimeout"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
}

type Mongo struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	URL        string `yaml:"url"`
	DBName     string `yaml:"db"`
	Collection string `yaml:"collection"`
}
