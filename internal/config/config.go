package config

import (
	"fmt"
	"marketflow/pkg/utils"
	"os"
	"time"
)

type Config struct {
	Postgres         PostgresConfig
	Redis            RedisConfig
	Exchanges        []Exchange
	PortAPI          int
	AggregatorWindow time.Duration
	RedisTTL         time.Duration
	AppEnv           string
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
	Port     int
	Name     string
	Password string
	DB       int
}

type Exchange struct {
	Name string
	Addr string
}

func Load() (*Config, error) {
	env := map[string]string{
		"APP_ENV":           os.Getenv("APP_ENV"),
		"PG_HOST":           os.Getenv("PG_HOST"),
		"PG_PORT":           os.Getenv("PG_PORT"),
		"PG_USER":           os.Getenv("PG_USER"),
		"PG_PASSWORD":       os.Getenv("PG_PASSWORD"),
		"PG_DB":             os.Getenv("PG_DB"),
		"PG_SSLMODE":        os.Getenv("PG_SSLMODE"),
		"REDIS_HOST":        os.Getenv("REDIS_HOST"),
		"REDIS_PORT":        os.Getenv("REDIS_PORT"),
		"REDIS_DB":          os.Getenv("REDIS_DB"),
		"EXCHANGE1_NAME":    os.Getenv("EXCHANGE1_NAME"),
		"EXCHANGE2_NAME":    os.Getenv("EXCHANGE2_NAME"),
		"EXCHANGE3_NAME":    os.Getenv("EXCHANGE3_NAME"),
		"EXCHANGE1_PORT":    os.Getenv("EXCHANGE1_PORT"),
		"EXCHANGE2_PORT":    os.Getenv("EXCHANGE2_PORT"),
		"EXCHANGE3_PORT":    os.Getenv("EXCHANGE3_PORT"),
		"EXCHANGE1_ADDR":    os.Getenv("EXCHANGE1_ADDR"),
		"EXCHANGE2_ADDR":    os.Getenv("EXCHANGE2_ADDR"),
		"EXCHANGE3_ADDR":    os.Getenv("EXCHANGE3_ADDR"),
		"API_PORT":          os.Getenv("API_PORT"),
		"AGGREGATOR_WINDOW": os.Getenv("AGGREGATOR_WINDOW"),
		"REDIS_TTL":         os.Getenv("REDIS_TTL"),
	}
	for key, value := range env {
		if value == "" {
			return nil, fmt.Errorf("missing required env variable: %s", key)
		}
	}
	pgPort, err := utils.ParseEnvInt("PG_PORT")
	if err != nil {
		return nil, err
	}

	redisPort, err := utils.ParseEnvInt("REDIS_PORT")
	if err != nil {
		return nil, err
	}

	redisDB, err := utils.ParseEnvInt("REDIS_DB")
	if err != nil {
		return nil, err
	}

	for i := 1; i <= 3; i++ {
		_, err := utils.ParseEnvInt("EXCHANGE%d_PORT")
		if err != nil {
			return nil, err
		}
	}

	portAPI, err := utils.ParseEnvInt("API_PORT")
	if err != nil {
		return nil, err
	}

	aggregatorWindow, err := utils.ValidTime("AGGREGATOR_WINDOW")
	if err != nil {
		return nil, err
	}

	redisTTL, err := utils.ValidTime("REDIS_TTL")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Postgres: PostgresConfig{
			Host:     os.Getenv("PG_HOST"),
			Port:     pgPort,
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			NameDB:   os.Getenv("PG_DB"),
			SSLMode:  os.Getenv("PG_SSLMODE"),
		},
		Redis: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     redisPort,
			Name:     os.Getenv("REDIS_DB"),
			Password: os.Getenv("PG_PASSWORD"),
			DB:       redisDB,
		},
		Exchanges: []Exchange{
			{
				Name: os.Getenv("EXCHANGE1_NAME"),
				Addr: os.Getenv("EXCHANGE2_ADDR"),
			},
			{
				Name: os.Getenv("EXCHANGE2_NAME"),
				Addr: os.Getenv("EXCHANGE2_ADDR"),
			},
			{
				Name: os.Getenv("EXCHANGE3_NAME"),
				Addr: os.Getenv("EXCHANGE3_ADDR"),
			},
		},
		PortAPI:          portAPI,
		AggregatorWindow: aggregatorWindow,
		RedisTTL:         redisTTL,
		AppEnv:           os.Getenv("APP_ENV"),
	}

	return cfg, nil
}
