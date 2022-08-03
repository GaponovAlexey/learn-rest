package app

import (
	"net/http"

	"github.com/GaponovAlexey/learn-rest/logging"
	"github.com/GaponovAlexey/learn-rest/pkg/app/config"
	"github.com/julienschmidt/httprouter"
	httpSwager "github.com/swaggo/http-swagger"
)

type App struct {
	config *config.Config
	logger *logging.Logger
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Println("router init")
	router := httprouter.New()

	logger.Println("swager")
	router.Handler(http.MethodGet,"/swagger", http.RedirectHandler("/swagger/inxex.html", http.StatusMovedPermanently) )
	router.Handler(http.MethodGet,"/swagger/*any", httpSwager.WrapHandler)
	return App{
		config: config,
		logger: logger,
	}, nil
}
