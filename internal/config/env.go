package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	JWTExpirationInSeconds int64
	JWTSecret              string
}

// create a singleton
var Envs = initConfig()

func initConfig() Config {
	godotenv.Load() //load environment variables from .env file
	return Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "orders_management"),
	}
}

func getEnv(key, fallback string) string {
	//Look for environment variable by key
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback // Default
}
