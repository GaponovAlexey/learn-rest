package main

import (
	"github.com/GaponovAlexey/learn-rest/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	s := server.New()
	if err := s.Start(); err != nil {
		logrus.Error(err)
	}
}
