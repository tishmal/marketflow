package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Postgres         PostgresConfig
	Redis            RedisConfig
	Exchanges        []Exchange
	AddrAPI          string
	AggregatorWindow time.Duration
	RedisTTL         time.Duration
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	NameDB   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Name     string
	Password string
	DB       int
}

type Exchange struct {
	Name string
	Port string
}

func Load() (*Config, error) {
	env := map[string]string{
		"PG_HOST":           os.Getenv("PG_HOST"),
		"PG_PORT":           os.Getenv("PG_PORT"),
		"PG_USER":           os.Getenv("PG_USER"),
		"PG_PASSWORD":       os.Getenv("PG_PASSWORD"),
		"PG_DB":             os.Getenv("PG_DB"),
		"PG_SSLMODE":        os.Getenv("PG_SSLMODE"),
		"REDIS_HOST":        os.Getenv("REDIS_HOST"),
		"REDIS_PORT":        os.Getenv("REDIS_PORT"),
		"REDIS_DB":          os.Getenv("REDIS_DB"),
		"EXCHANGE1_ADDR":    os.Getenv("EXCHANGE1"),
		"EXCHANGE2_ADDR":    os.Getenv("EXCHANGE2_ADDR"),
		"EXCHANGE3_ADDR":    os.Getenv("EXCHANGE3_ADDR"),
		"API_ADDR":          os.Getenv("API_ADDR"),
		"AGGREGATOR_WINDOW": os.Getenv("AGGREGATOR_WINDOW"),
		"REDIS_TTL":         os.Getenv("REDIS_TTL"),
	}
	for key, value := range env {
		if value == "" {
			return nil, fmt.Errorf("missing required env variable: %s", key)
		}
	}

	pgPort, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		return nil, fmt.Errorf("error getting postgres port :%s", err)
	}

	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return nil, fmt.Errorf("error getting redis port :%s", err)
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, fmt.Errorf("error getting redis db :%s", err)
	}

	for i := 1; i <= 3; i++ {
		portStr := os.Getenv(fmt.Sprintf("EXCHANGE%d_PORT", i))
		_, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("error getting redis db: %s", err)
		}
	}

	cfg := &Config{
		Postgres: PostgresConfig{
			Host: os.Getenv("PG_HOST"),
			Port: pgPort,
		},
	}

	return &cfg, nil
}
