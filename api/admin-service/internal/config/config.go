package config

import "os"

type Config struct {
	AdminPort  string
	ServerHost string
	ProductAPI string
	UserAPI    string
	CartAPI    string
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadConfig() Config {
	return Config{
		AdminPort:  getEnv("PORT", "8084"),
		ServerHost: getEnv("SERVER_HOST", "localhost"),
		ProductAPI: getEnv("PRODUCT_API_URL", "http://product-service:8082"),
		UserAPI:    getEnv("USER_API_URL", "http://user-service:8081"),
		CartAPI:    getEnv("CART_API_URL", "http://cart-service:8083"),
	}
}
