package main

import (
	"log"

	"github.com/GaponovAlexey/learn-rest/logging"
	"github.com/GaponovAlexey/learn-rest/pkg/app/config"
	"github.com/GaponovAlexey/learn-rest/pkg/app"
)

func main() {
	log.Println("config")
	cfg := config.GetConfig()

	log.Println("logger")
	l := logging.GetLogger()

	_, err := app.NewApp(cfg, l)
	if err != nil {
		l.Fatal()
	}
}
