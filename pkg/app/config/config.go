package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	isDebug       bool `env:"IS_DEBUG" env-default:"false"`
	isDevelopment bool `env:"IS_DEV" env-default:"false"`
	// listen        struct {
	Type   string `env:"LISTEN_TYPE" env-default:"port"`
	BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
	port   string `env:"PORT" env-default:"3000"`
	// }
	// AppConfig struct {
	LogLevel string `env:"PORT" env-default:"3000"`
	// AdminUser struct {
	Email    string `env:"ADMIN_EMAIL" env-default:"true"`
	Password string `env:"ADMIN_PWD" env-default:"true"`
	// }
	// }
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("config")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			var helpText = "help"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Println(help)
			log.Println(err)
		}
	})
	return instance
}
