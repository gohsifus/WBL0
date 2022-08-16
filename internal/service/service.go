package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"natTest/internal/apiserver"
	"natTest/internal/natsSubscriber"
	"natTest/internal/store"
)

//Service структура описыващая сервис
type Service struct {
	Configs   *Config
	ApiServer *apiserver.APIServer
	Nats      *natsSubscriber.NatsSubscriber
	Store     *store.Store
	Logger    *logrus.Logger
}

//New создаст сервис
func New(configs *Config) (*Service, error) {
	logger := logrus.New()
	store := store.New(configs.StoreConfig)
	nats, err := natsSubscriber.New(configs.NatsConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к nats: %s", err)
	}

	return &Service{
		Configs:   configs,
		ApiServer: apiserver.New(configs.ApiServerConfig, logger),
		Nats:      nats,
		Store:     store,
		Logger:    logger,
	}, nil
}

func (s *Service) Start() error {
	s.Logger.Info("Запуск сервиса")

	err := s.configureLogger()
	if err != nil {
		return err
	}
	s.Logger.Info("Настройка логирования")

	err = s.configureStore()
	if err != nil {
		return err
	}
	s.Logger.Info("Настройка подключения к store")

	err = s.ApiServer.Start()
	if err != nil {
		return err
	}
	s.Logger.Info("Запуск APIServer")

	return nil
}

//configureLogger настроит логирование в сервисе
func (s *Service) configureLogger() error {
	level, err := logrus.ParseLevel(s.Configs.LogLevel)
	if err != nil {
		return fmt.Errorf("ошибка настройки логирования: %s", err)
	}

	s.Logger.SetLevel(level)
	return nil
}

//configureStore настроит подключение к store
func (s *Service) configureStore() error {
	err := s.Store.Open()
	if err != nil {
		return fmt.Errorf("ошибка нстройки хранилища: %s", err)
	}

	return nil
}
