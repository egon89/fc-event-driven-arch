package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	KafkaBroker      string
	KafkaTopic       string
	AppPort          string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseHost:     getEnv("DB_HOST", ""),
		DatabasePort:     getEnv("DB_PORT", ""),
		DatabaseName:     getEnv("DB_NAME", ""),
		DatabaseUser:     getEnv("DB_USER", ""),
		DatabasePassword: getEnv("DB_PASSWORD", ""),
		KafkaBroker:      getEnv("KAFKA_BROKER", "kafka:9092"),
		KafkaTopic:       getEnv("KAFKA_TOPIC", "balances"),
		AppPort:          getEnv("APP_PORT", "8080"),
	}

	if cfg.DatabaseHost == "" {
		log.Fatal("DB_HOST is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
