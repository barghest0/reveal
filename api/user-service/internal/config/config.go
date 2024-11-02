package config

import "os"

type Config struct {
	Port        string
	ServerHost  string
	RabbitMQURL string
}

type DBConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
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
	AppConfig.RabbitMQURL = "amqp://guest:guest@localhost:5672/"

	return AppConfig
}

func LoadDBConfig() DBConfig {
	var DBConfig DBConfig

	DBConfig.User = getEnv("POSTGRES_USER", "postgres")
	DBConfig.Password = getEnv("POSTGRES_PASSWORD", "barghest")
	DBConfig.Name = getEnv("POSTGRES_DB", "users")
	DBConfig.Host = getEnv("DB_HOST", "postgres")
	DBConfig.Port = getEnv("DB_PORT", "5432")

	return DBConfig
}
