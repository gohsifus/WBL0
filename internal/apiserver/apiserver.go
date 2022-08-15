package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"natTest/internal/store"
	"net/http"
)

type APIServer struct {
	config *Config
	Logger *logrus.Logger
	router *mux.Router
	Store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		Logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start запуск сервиса
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.Logger.Info("Запуск APIServer")

	http.ListenAndServe(s.config.BindAddr, s.router)

	return nil
}

//configureLogger сконфигурирует логирование
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.Logger.SetLevel(level)

	return nil
}

//configureStore иницмализирует подключение к бд
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	//Проверяем подключение
	if err := st.Open(); err != nil {
		return err
	}
	s.Logger.Info("Подключение к БД")
	//Инициализируем хранилище
	s.Store = st

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
