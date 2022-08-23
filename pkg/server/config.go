package server

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
		LogLevel: "debug",
	}
}
