package main

import (
	"github.com/GaponovAlexey/learn-rest/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	config := server.NewConfig()
	s := server.New(config)
	if err := s.Start(); err != nil {
		logrus.Error(err)
	}
}
