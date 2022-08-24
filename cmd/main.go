package main

import (
	"log"

	"github.com/GaponovAlexey/learn-rest/pkg/app/apiserver"
)

func main() {

	config := apiserver.NewConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
