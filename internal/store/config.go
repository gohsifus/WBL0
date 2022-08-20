package store

type Config struct {
	DatabaseUrl string `json:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
