package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"natTest/internal/natsSubscriber"
	"natTest/internal/store"
	"os"
)

type Config struct{
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"` //Уровень вывода
	Store *store.Config `json:"store"`
	Nats *natsSubscriber.Config `json:"nats"`
}

//LoadConfig Проинициализирует значения конфигурации
func (c *Config) LoadConfig(path string) error{
	file, err := os.Open(path)
	if err != nil{
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil{
		return err
	}

	err = json.Unmarshal(data, &c)
	if err != nil{
		return err
	}

	return nil
}

//NewConfig Вернет конфигурацию с значениями по умолчанию
func NewConfig() *Config{
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
		Nats: natsSubscriber.NewConfig(),
	}
}