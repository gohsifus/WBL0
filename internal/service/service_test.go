package service

import (
	"encoding/json"
	"os"
	"testing"
)

func TestConfig_LoadConfig(t *testing.T) {
	config := NewConfig()
	testPath := "./test.json"

	//Создаем временный файл конфигураций
	tmpFile, err := os.Create(testPath)
	defer os.Remove(testPath)
	defer tmpFile.Close()
	if err != nil {
		t.Skip(err)
	}

	configs := `
	{
	  "apiServerConfig": {
		"bind_addr": ":8080"
	  },
	  "natsConfig": {
		"cluster_id": "test-cluster",
		"client_id": "test-sub",
		"channel_name": "orders"
	  },
	  "storeConfig": {
		"database_url": "user=testuser password=qawsed345rf dbname=dbForL0 sslmode=disable"
	  },
	  "log_level": "debug"
	}`
	//Записываем в временный файл
	_, err = tmpFile.Write([]byte(configs))
	if err != nil {
		t.Skip(err)
	}

	err = config.LoadConfig(testPath)
	if err != nil {
		t.Errorf("возникла неожиданная ошибка: %v", err)
	}

	expected := &Config{}
	err = json.Unmarshal([]byte(configs), &expected)
	if err != nil {
		t.Skip(err)
	}

	if  *expected.ApiServerConfig != *config.ApiServerConfig ||
		*expected.NatsConfig != *config.NatsConfig ||
		*expected.StoreConfig != *config.StoreConfig ||
		expected.LogLevel != config.LogLevel {
			t.Errorf("ожидаемая конфигурация не совпадает с фактической")
	}
}
