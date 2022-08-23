package server

type Config struct {
	BindAddr string `json:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
	}
}
