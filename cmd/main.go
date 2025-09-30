package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/mrbelka12000/pdd_tests_bot/internal/client/ai"
	"github.com/mrbelka12000/pdd_tests_bot/internal/delivery"
	"github.com/mrbelka12000/pdd_tests_bot/internal/repo"
	"github.com/mrbelka12000/pdd_tests_bot/internal/usecase"
	"github.com/mrbelka12000/pdd_tests_bot/pkg/config"
	"github.com/mrbelka12000/pdd_tests_bot/pkg/gorm/postgres"
	"github.com/mrbelka12000/pdd_tests_bot/pkg/redis"
	"github.com/mrbelka12000/pdd_tests_bot/pkg/storage/minio"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service_name", cfg.ServiceName)

	db, err := postgres.New(cfg.PGURL)
	if err != nil {
		log.With("error", err).Error("failed to connect to database")
		return
	}

	if err := runMigrations(cfg.PGURL); err != nil {
		log.With("error", err).Error("failed to run migrations")
		return
	}

	repo := repo.New()

	rCache, err := redis.New(cfg)
	if err != nil {
		log.With("error", err).Error("failed to connect to redis")
		return
	}

	minIOClient, err := minio.Connect(cfg)
	if err != nil {
		log.With("error", err).Error("failed to connect to minio")
		return
	}

	uc := usecase.New(db, repo.User, repo.Case, minIOClient, ai.NewClient(cfg.AIToken))

	if cfg.NeedToImportData {
		if err := uc.Import(); err != nil {
			log.With("error", err).Error("failed to import data")
			return
		}
	}

	_, err = delivery.Start(cfg, uc, log, rCache)
	if err != nil {
		log.With("error", err).Error("failed to start delivery")
		return
	}
}

func runMigrations(dsn string) error {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open postgres to migration, err: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect postgres to migration, err: %w", err)
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		return fmt.Errorf("failed to run postgres migration, err: %w", err)
	}

	return nil
}
