package natsSubscriber

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"natTest/pkg/models"
	"sync"
)

type NatsSubscriber struct{
	conn stan.Conn
	config *Config
}

//New Произведет подключение к nats и вернет структуру с подключением
func New(configs *Config) (*NatsSubscriber, error){
	conn, err := stan.Connect(configs.ClusterId, configs.ClientId)
	if err != nil{
		return nil, err
	}
	//logger.Info("Подключение к NATS")

	return &NatsSubscriber{
		conn: conn,
		config: configs,
	}, nil
}

func (n *NatsSubscriber) GetDataFromChannel() error{
	_, err := n.conn.Subscribe("orders", func(m *stan.Msg) {
		o := models.Order{}
		json.Unmarshal(m.Data, &o)
		fmt.Printf("%+v\n", o)
		//fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil{
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	return nil
}

