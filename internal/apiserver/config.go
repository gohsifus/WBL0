package apiserver

type Config struct {
	BindAddr string `json:"bind_addr"`
}

//NewConfig Вернет конфигурацию с значениями по умолчанию
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
