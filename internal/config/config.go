package config

import (
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"pos-receipt-system"`
	Host        string `envconfig:"HOST" default:"localhost"`
	Port        string `envconfig:"PORT" default:"8080"`
	DatabaseURL string `envconfig:"DATABASE_URL" default:"postgres://user:password@localhost:5432/pos_receipt_system"`
	Logger      *zap.SugaredLogger
}

func NewConfig(ctx context.Context) (Config, error) {
	c := NewDefault()
	if err := envconfig.Process("", &c); err != nil {
		return Config{}, fmt.Errorf("failed to process environment variables: %w", err)
	}
	if err := c.initLogger(); err != nil {
		return Config{}, fmt.Errorf("failed to initialize logger: %w", err)
	}
	return c, nil
}

func NewDefault() Config {
	return Config{
		ServiceName: "pos-receipt-system",
		Host:        "localhost",
		Port:        "8080",
		DatabaseURL: "postgres://user:password@localhost:5432/pos_receipt_system",
		Logger:      zap.NewExample().Sugar(),
	}
}

func (c *Config) initLogger() error {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // 👈 set log level to debug

	logger, err := cfg.Build()
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}

	c.Logger = logger.Sugar()
	return nil
}
