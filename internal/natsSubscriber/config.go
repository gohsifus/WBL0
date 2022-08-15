package natsSubscriber

type Config struct{
	ChannelName string `json:"channel_name"`
	ClusterId string `json:"cluster_id"`
	ClientId string `json:"client_id"`
}

func NewConfig() *Config{
	return &Config{}
}
