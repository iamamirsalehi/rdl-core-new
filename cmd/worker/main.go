package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rdl/core/internal/platform/config"
	"github.com/rdl/core/internal/platform/logger"
	"github.com/rdl/core/internal/platform/mongodb"
	"github.com/rdl/core/internal/platform/nats"
	"github.com/rdl/core/internal/platform/postgres"
	"github.com/rdl/core/internal/platform/redis"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log := logger.New(cfg.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongo, err := mongodb.Connect(ctx, cfg.MongoDB)
	if err != nil {
		log.Error("failed to connect to mongodb", "error", err)
		os.Exit(1)
	}
	defer mongo.Disconnect(ctx)

	pg, err := postgres.Connect(cfg.Postgres)
	if err != nil {
		log.Error("failed to connect to postgres", "error", err)
		os.Exit(1)
	}
	defer pg.Close()

	rdb, err := redis.Connect(cfg.Redis)
	if err != nil {
		log.Error("failed to connect to redis", "error", err)
		os.Exit(1)
	}
	defer rdb.Close()

	nc, err := nats.Connect(cfg.NATS)
	if err != nil {
		log.Error("failed to connect to nats", "error", err)
		os.Exit(1)
	}
	defer nc.Close()

	log.Info("worker started, waiting for messages...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down worker...")
}
