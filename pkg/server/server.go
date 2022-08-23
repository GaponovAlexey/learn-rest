package server

import (
	"io"
	"net/http"

	"github.com/GaponovAlexey/learn-rest/pkg/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//start
func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("start api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//logger
func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

//router
func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/", s.Hallo())
}

// func
func (s *ApiServer) Hallo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hi")
	}
}
func (s *ApiServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
