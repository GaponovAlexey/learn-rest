package main

import (
	"log"

	"github.com/GaponovAlexey/learn-rest/logging"
	"github.com/GaponovAlexey/learn-rest/pkg/app/config"
)

func main() {
	log.Println("config")
	cfg := config.GetConfig()

	log.Println("logger")
	l := logging.GetLogger()

	app, err := app.NewApp(cfg, l)
	if err != nil {
		l.Fatal()
	}
}
