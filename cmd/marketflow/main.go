package main

import (
	"flag"
	"log"
	"marketflow/internal/config"
	"marketflow/pkg/logger"
)

func main() {
	portFlag := flag.Int("port", 8080, "port number")
	help := flag.Bool("help", false, "show usage")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}
	//...
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// Переопределение порта, если передан через флаг
	if *portFlag != 0 {
		cfg.PortAPI = *portFlag
	}

	logger.Init(cfg.AppEnv)
	logger.Info("Logger initialized", "env", cfg.AppEnv)

	logger.Debug("Loaded config", "config", cfg)

	// initial interfaces
	logger.Info("Starting marketflow...", "port", cfg.PortAPI)
}
