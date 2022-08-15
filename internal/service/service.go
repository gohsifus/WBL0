package service

import (
	"github.com/sirupsen/logrus"
	"natTest/internal/apiserver"
	"natTest/internal/natsSubscriber"
	"natTest/internal/store"
)

//Service структура описываюзая сервис
type Service struct {
	ApiServer *apiserver.APIServer
	Nats      *natsSubscriber.NatsSubscriber
	Store     *store.Store
	Logger    *logrus.Logger
}

func New() *Service{
	return &Service{}
}
