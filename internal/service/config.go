package service

import (
	"encoding/json"
	"io/ioutil"
	"natTest/internal/apiserver"
	"natTest/internal/natsSubscriber"
	"natTest/internal/store"
	"os"
)

//Config конфигурации приложения
type Config struct {
	apiServerConfig *apiserver.Config
	natsConfig      *natsSubscriber.Config
	storeConfig     *store.Config
	logLevel        string
}

//NewConfig вернет конфигурацию заполненную значениями по умолчанию
func NewConfig() *Config {
	return &Config{
		apiServerConfig: apiserver.NewConfig(),
		natsConfig:      natsSubscriber.NewConfig(),
		storeConfig:     store.NewConfig(),
		logLevel:        "debug",
	}
}

//LoadConfig инициализирует конфигурацию из файла path
func (c *Config) LoadConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	return nil
}
