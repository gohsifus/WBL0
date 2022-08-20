// Package apiserver http Сервер отображающий данные из кеша
package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"html/template"
	"natTest/internal/cache"
	"net/http"
	"strconv"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	cache  cache.Cache
	templateCache *template.Template
}

func New(config *Config, logger *logrus.Logger, cache cache.Cache) *APIServer {
	return &APIServer{
		config: config,
		logger: logger,
		router: mux.NewRouter(),
		cache:  cache,
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

//configureRouter Назначает функции обработчики на url
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/hello", s.handleHello())

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/"))))
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}
}

//handleIndex обработчик главной страницы
func (s *APIServer) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if s.templateCache == nil {
			files := []string{
				"./web/html/index.html",
				"./web/html/templates/header.html",
				"./web/html/templates/footer.html",
			}

			temp, err := template.ParseFiles(files...)
			if err != nil {
				s.logger.Error(err)
			}

			s.templateCache = temp
		}

		var dataForPage interface{}
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err == nil {
			dataForPage = s.cache.GetById(id)
		} else {
			dataForPage = nil
		}

		err = s.templateCache.Execute(w, dataForPage)
		if err != nil {
			s.logger.Error(err)
		}
	}
}
