package natsSubscriber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"natTest/pkg/models"
)

//NatsSubscriber описывает подключение к nats
type NatsSubscriber struct {
	config *Config
	conn   stan.Conn
	logger *logrus.Logger
	validator *validator.Validate
}

//New Произведет подключение к nats и вернет структуру с подключением
func New(configs *Config, logger *logrus.Logger) (*NatsSubscriber, error) {
	conn, err := stan.Connect(configs.ClusterId, configs.ClientId)
	if err != nil {
		return nil, err
	}

	return &NatsSubscriber{
		conn:   conn,
		config: configs,
		logger: logger,
		validator: validator.New(),
	}, nil
}

//GetDataFromChannel подпишется на канал и будет возвращать данные из него
func (n *NatsSubscriber) GetDataFromChannel(channelName string) (<-chan models.Order, error) {
	out := make(chan models.Order)

	_, err := n.conn.Subscribe(channelName, func(m *stan.Msg) {
		recOrder := models.Order{}

		decoder := json.NewDecoder(bytes.NewReader(m.Data))
		decoder.DisallowUnknownFields() //Бракуем данные если в json лишние поля

		err := decoder.Decode(&recOrder)
		err = n.validator.Struct(recOrder)

		if err != nil {
			//Игнорируем некорректные данные
			n.logger.Info("ignore bad data")
		} else {
			out <- recOrder
		}
	}, stan.DurableName("durableId"), stan.StartWithLastReceived())

	if err != nil {
		return nil, fmt.Errorf("ошибка оформления подписки на канал %s: %s", channelName, err)
	}

	return out, nil
}
