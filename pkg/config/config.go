package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type (
	// Config of service
	Config struct {
		InstanceConfig
		DBConfig
		ClientsConfig
		RedisConfig
		MinIOConfig
		TelegramConfig
	}

	InstanceConfig struct {
		ServiceName      string `env:"SERVICE_NAME,required"`
		NeedToImportData bool   `env:"NEED_TO_IMPORT_DATA,default=false"`
	}

	DBConfig struct {
		PGURL          string `env:"PG_URL,required"`
		MigrationsPath string `env:"MIGRATIONS_PATH, default=migrations/"`
	}

	ClientsConfig struct {
		AIToken string `env:"AI_TOKEN,required"`
	}

	RedisConfig struct {
		RedisAddr string `env:"REDIS_ADDR,required"`
	}

	MinIOConfig struct {
		MinIOAddr      string `env:"MINIO_ADDR,required"`
		MinIOBucket    string `env:"MINIO_BUCKET,default=linguo_sphere"`
		MinIOAccessKey string `env:"MINIO_ACCESS_KEY,required"`
		MinIOSecretKey string `env:"MINIO_SECRET_KEY,required"`
	}

	TelegramConfig struct {
		BotToken string `env:"BOT_TOKEN,required"`
	}
)

// Get
func Get() (Config, error) {
	return parseConfig()
}

func parseConfig() (cfg Config, err error) {
	godotenv.Load()

	err = envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return cfg, fmt.Errorf("fill config: %w", err)
	}

	return cfg, nil
}
