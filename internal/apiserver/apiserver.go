package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config, logger *logrus.Logger) *APIServer {
	return &APIServer{
		config: config,
		logger: logger,
		router: mux.NewRouter(),
	}
}

//Start запуск сервера
func (s *APIServer) Start() error {
	s.configureRouter()

	err := http.ListenAndServe(s.config.BindAddr, s.router)
	if err != nil {
		fmt.Errorf("ошибка запуска APIServer: %s", err)
	}

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}
}
