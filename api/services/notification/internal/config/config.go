package config

import "os"

type Config struct {
	Port       string
	ServerHost string
	Exchange   string
	OrderQueue string
	BrokerURL  string
	Token      string
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadConfig() Config {
	var AppConfig Config

	AppConfig.Port = getEnv("PORT", "8080")
	AppConfig.ServerHost = getEnv("SERVER_HOST", "localhost")
	AppConfig.Exchange = "notifications"
	AppConfig.OrderQueue = "orders"
	AppConfig.BrokerURL = "amqp://guest:guest@rabbitmq:5672/"
	AppConfig.Token = "7782679125:AAG33kQ2c7DX0AfUAyvgMFRjWI2mIuzKy7c"

	return AppConfig
}
