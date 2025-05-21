package main

import (
	"context"
	"flag"
	"log"
	"marketflow/internal/adapter/cache"
	"marketflow/internal/adapter/storage"
	"marketflow/internal/app/port/out"
	"marketflow/internal/config"
	"marketflow/pkg/logger"
	"time"
)

func main() {
	// флаги
	portFlag := flag.Int("port", 8080, "port number")
	help := flag.Bool("help", false, "show usage")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}
	// контекст
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// конфиги
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// если порт задан
	if *portFlag != 0 {
		cfg.PortAPI = *portFlag
	}

	logger.Init(cfg.AppEnv)
	logger.Info("Logger initialized", "env", cfg.AppEnv)

	logger.Debug("Loaded config", "config", cfg)
	// подключение бд
	repo, err := storage.NewPostgresRepository(ctx, cfg.Postgres)
	if err != nil {
		log.Fatal("failed to init postgres: %w", err)
	}
	defer repo.Close()

	var cachePort out.CachePort = cache.NewRedisCache(cfg.Redis, cfg.RedisTTL)

	if err := cachePort.ConnectCache(ctx); err != nil {
		logger.Error("Redis connection failed", "err", err)
		return
	}

	logger.Info("Starting marketflow...", "port", cfg.PortAPI)
}
