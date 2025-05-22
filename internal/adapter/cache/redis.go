package cache

import (
	"context"
	"fmt"
	"marketflow/internal/config"
	"marketflow/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client      *redis.Client
	config      config.RedisConfig
	ttl         time.Duration
	isConnected bool
}

// создает новый экземпляр RedisCache
func NewRedisCache(cfg config.RedisConfig, ttl time.Duration) *RedisCache {
	return &RedisCache{
		config: cfg,
		ttl:    ttl,
	}
}

// устанавливает соединение с Redis
func (r *RedisCache) ConnectCache(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", r.config.Host, r.config.Port)

	r.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.config.Password,
		DB:       r.config.DB,
	})

	if err := r.client.Ping(ctx).Err(); err != nil {
		logger.Error("Failed to connect to Redis", "error", err)
		return fmt.Errorf("failed to ping redis: %w", err)
	}

	r.isConnected = true
	logger.Info("Successfuly connecnted to Redis", "host", r.config.Host, r.config.Port)

	return nil
}

func (r *RedisCache) Close() error {
	if r.client == nil {
		return nil
	}

	err := r.client.Close()
	if err != nil {
		logger.Error("Failed to close Redis connection", "error", err)
		return fmt.Errorf("failed to close Redis connection: %w", err)
	}

	r.isConnected = false
	logger.Info("Redis connection closed")
	return nil
}
