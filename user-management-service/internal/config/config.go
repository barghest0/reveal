package config

import "os"

type Config struct {
	Port          string
	ServerAddress string
	DBConfig      DBConfig
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

var AppConfig Config

func LoadConfig() Config {
	AppConfig.Port = getEnv("PORT", "8080")
	AppConfig.ServerAddress = getEnv("SERVER_ADDRESS", "localhost")
	AppConfig.DBConfig.User = getEnv("POSTGRES_USER", "postgres")
	AppConfig.DBConfig.Password = getEnv("POSTGRES_PASSWORD", "barghest")
	AppConfig.DBConfig.Name = getEnv("POSTGRES_DB", "users")
	AppConfig.DBConfig.Host = getEnv("DB_HOST", "localhost")
	AppConfig.DBConfig.Port = getEnv("DB_PORT", "5432")

	return AppConfig
}
