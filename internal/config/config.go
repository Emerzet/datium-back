package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	App  AppConfig
	HTTP HTTPConfig
	DB   DbConfig
}

type AppConfig struct {
	Env string
	BaseURL string
}

type HTTPConfig struct {
	Addr        string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	MaxBodyBytes int64

}

type DbConfig struct {
	Host string
	Port int
	Database string
	User string
	Password string



}







func Load() Config {
	return Config{
		App: AppConfig {
			Env:"dev",
			BaseURL: "http://localhost:8080",
		},
	
		HTTP: HTTPConfig {
			Addr:"localhost:8080",
			ReadTimeout: 30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout: 120 * time.Second,	
			MaxBodyBytes: 1 << 20,
		},
		
		DB: DbConfig {
			Host: os.Getenv("DB_HOST"),
			Port: getEnvAsInt("DB_PORT", 5432),
			Database: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),

		},
					}

	}

	func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val

	}




