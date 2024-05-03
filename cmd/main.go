package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/config"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/errors"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/http/server"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/log"
	repo "github.com/nick1729/resp-api-tmpl/internal/pkg/repository"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/storage/postgres"
)

const (
	appName  = "resp-api-tmpl"
	timezone = "Europe/Moscow"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return errors.Wrap(err, "loading location")
	}

	time.Local = location

	cfg, err := config.Init()
	if err != nil {
		return errors.Wrap(err, "initializing config")
	}

	logger := log.Init(cfg.Log)

	ctx := context.TODO()

	pgPool, err := postgres.New(ctx, cfg.Postgres, logger)
	if err != nil {
		return errors.Wrap(err, "creating pg pool")
	}
	defer pgPool.Close()

	repo := repo.New(pgPool)

	s := server.New(cfg.Server, appName, repo)
	address := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	logger.Info().Msgf("Starting HTTP server on %s", address)

	s.Run(address)

	sigHandler := make(chan os.Signal, 1)
	signal.Notify(sigHandler, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-sigHandler:
		logger.Info().Msg("Gracefully shutdown")
	case err := <-s.Notify():
		logger.Error().Err(err).Msg("Starting HTTP server")
	}

	s.Stop()

	return nil
}
