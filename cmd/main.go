package main

import (
	"fmt"

	"github.com/GaponovAlexey/learn-rest/pkg/app/config"
)

func main() {

	err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
	}

}
