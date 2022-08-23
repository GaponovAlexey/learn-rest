package server

import "github.com/GaponovAlexey/learn-rest/pkg/store"

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
