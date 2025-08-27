package app

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment variables.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisHost string
	RedisPort string

	APIPort string

	JWTSecret       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

// LoadConfig reads configuration from the environment and optional .env file.
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	accessTTL, _ := time.ParseDuration(os.Getenv("JWT_ACCESS_TTL"))
	refreshTTL, _ := time.ParseDuration(os.Getenv("JWT_REFRESH_TTL"))

	cfg := &Config{
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		RedisHost:       os.Getenv("REDIS_HOST"),
		RedisPort:       os.Getenv("REDIS_PORT"),
		APIPort:         os.Getenv("API_PORT"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		AccessTokenTTL:  accessTTL,
		RefreshTokenTTL: refreshTTL,
	}

	return cfg, nil
}
