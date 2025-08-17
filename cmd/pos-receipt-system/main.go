package main

import (
	"context"

	"github.com/nitinjangam/pos-receipt-system/api"
	"github.com/nitinjangam/pos-receipt-system/internal/config"
	"github.com/nitinjangam/pos-receipt-system/internal/db"
	"github.com/nitinjangam/pos-receipt-system/internal/handler"
	"github.com/nitinjangam/pos-receipt-system/internal/repository"
	"github.com/nitinjangam/pos-receipt-system/internal/service"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("pos-receipt-system")

const (
	ServiceName  = "pos-receipt-system"
	DatabasePath = "./pos-receipt-system.db"
)

func main() {
	ctx := context.Background()

	// Initialize config
	config, err := config.NewConfig(ctx)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	// Initialize database - dirty way ToDo: use a proper database in future
	db := db.InitSQLite(DatabasePath)
	// Initialize services
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, config.Logger)
	authHandler := handler.NewAuthHandler(authService, config.Logger)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, config.Logger)
	productHandler := handler.NewProductHandler(productService, config.Logger)

	salesRepository := repository.NewSalesRepository(ctx)
	salesService := service.NewSalesService(tracer, config.Logger, salesRepository)
	salesHandler := handler.NewSalesHandler(ctx, config.Logger, salesService)

	settingsRepository := repository.NewSettingsRepository(ctx)
	settingsService := service.NewSettingsService(tracer, config.Logger, settingsRepository)
	settingsHandler := handler.NewSettingsHandler(tracer, config.Logger, settingsService)

	// ToDo: create health check service

	handler := handler.NewHandler(authHandler, productHandler, salesHandler, settingsHandler)

	// Run the API
	if err := api.Run(ctx, config, handler); err != nil {
		config.Logger.Errorw("failed to run API", "error", err)
		panic("failed to run API: " + err.Error())
	} else {
		config.Logger.Infow("API is running", "port", config.Port)
	}
}
