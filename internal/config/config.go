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
	Env       string
	BaseURL   string
	StaticDir string
}

type HTTPConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	MaxHeaderBytes int
	MaxBodyBytes int64
}

type DbConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func Load() Config {
	return Config{
		App: AppConfig{
			Env:       getEnv("APP_ENV", "dev"),
			BaseURL:   getEnv("BASE_URL", "http://localhost:8080"),
			StaticDir: getEnv("STATIC_DIR", "/Users/mariuszdubicki/Documents/dev/datium-front/public"),
		},

		HTTP: HTTPConfig{
			Addr:         getEnv("HTTP_ADDR", ":8080"),
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
			MaxHeaderBytes: 1 << 20,
			MaxBodyBytes: 1 << 20,
		},

		DB: DbConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			Database: getEnv("DB_NAME", "datium"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
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

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val

}


