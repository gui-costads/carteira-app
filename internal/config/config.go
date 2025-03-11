package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfig struct {
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	JWTSecretKey  []byte
	JWTExpiration time.Duration
	ServerPort    string
	Environment   string
}

func Load() *AppConfig {
	_ = godotenv.Load()

	cfg := &AppConfig{
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnvInt("DB_PORT", 5432),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", ""),
		DBName:        getEnv("DB_NAME", "carteira"),
		DBSSLMode:     getEnv("DB_SSL_MODE", "disable"),
		JWTSecretKey:  []byte(getEnv("JWT_SECRET_KEY", "default-secret-unsafe")),
		JWTExpiration: parseDuration(getEnv("JWT_EXPIRATION", "24h")),
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		Environment:   getEnv("APP_ENV", "development"),
	}

	return cfg
}

func (c *AppConfig) DatabaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}
	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return defaultValue
}

func parseDuration(durationStr string) time.Duration {
	dur, err := time.ParseDuration(durationStr)
	if err != nil {
		return 24 * time.Hour
	}
	return dur
}
