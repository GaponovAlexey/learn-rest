package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initial config: %s", err.Error())
	}
	fmt.Println(viper.ReadConfig())

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
