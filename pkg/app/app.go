package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/GaponovAlexey/learn-rest/logging"
	"github.com/GaponovAlexey/learn-rest/pkg/app/config"
	"github.com/julienschmidt/httprouter"
	httpSwager "github.com/swaggo/http-swagger"
)

type App struct {
	cfg    *config.Config
	logger *logging.Logger
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Println("router init")
	router := httprouter.New()

	//swagger
	logger.Println("swager")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/inxex.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwager.WrapHandler)

	//return
	return App{
		cfg:    config,
		logger: logger,
	}, nil
}

func (a *App) startHttp() {
	a.logger.Info("start Http")
	
	var listener net.Listener

	if a.cfg.Type == "sock" {
		addDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			a.logger.Fatal(err)
		}
		socketPath := path.Join(addDir, a.cfg.SocketFile)
		a.logger.Infof("socketPath %s", socketPath)

		a.logger.Info("create and listen unix socket")
		listener, err := net.Listen("unix", socketPath)
		if err != nil {
			a.logger.Fatal(err)
		}
	} else {
		a.logger.Infof("bind apli%s and port%s", a.cfg.BindIP, a.cfg.Port)
		var err error
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.BindIP, a.cfg.Port))
		if err != nil {
			a.logger.Fatal(err)
		}
	}
}
