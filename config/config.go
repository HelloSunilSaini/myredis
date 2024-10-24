package config

import "time"

type Config struct {
	ServerPort               int
	MaxConnections           int
	MaxIdleConnectionTimeout time.Duration
	InputBufferSize          int
}

func GetConfig() *Config {
	return &Config{
		ServerPort:               6379,
		MaxConnections:           1000,
		MaxIdleConnectionTimeout: time.Second * 60,
		InputBufferSize:          1024,
	}
}
