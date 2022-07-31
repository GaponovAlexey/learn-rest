package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/GaponovAlexey/learn-rest/pkg/app/configToml"
)

func main() {
	flag.Parse()
	var configPath string
	fmt.Println("c:", configPath)

	conf := configToml.NewConfig()

	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(conf)

}
