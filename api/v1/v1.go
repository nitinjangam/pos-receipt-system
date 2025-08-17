package v1

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

//go:generate oapi-codegen -generate types,gin,spec -package v1 -o api-gen.go openapi.yaml

const (
	ServiceName   = "pos-receipt-system"
	SchemaVersion = "1.0"
)

type Config struct {
	Port     string
	Services []ServerInterface
	Logger   *zap.SugaredLogger
}

type Handler struct {
	config Config
	logger *zap.SugaredLogger
	srv    *http.Server
}

func New(config *Config) (*Handler, error) {
	swagger, err := GetSwagger()
	if err != nil {
		return nil, err
	}

	r := gin.New()

	// CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	swagger.Servers = nil // Skip servers validation as recommended by oapi-codegen package

	r.Use(otelgin.Middleware(ServiceName))

	options := &middleware.Options{}

	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, options))

	r.Use(gin.Recovery())

	l := config.Logger.Desugar()

	r.Use(ginzap.RecoveryWithZap(l, true))

	for _, val := range config.Services {
		RegisterHandlers(r, val)
	}

	httpServer := &http.Server{
		ReadTimeout: 5 * time.Minute, // TODO: verify this and parameterize
		Addr:        config.Port,
		Handler:     r,
	}

	return &Handler{
		srv:    httpServer,
		config: *config,
		logger: config.Logger,
	}, nil
}

func getRequestBody(c *gin.Context) string {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return ""
	}
	return string(payload)
}

func (h *Handler) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errWg, errCtx := errgroup.WithContext(ctx)
	h.logger.Info("API v1: starting server at", h.config.Port)
	// start HTTP server in one goroutine
	errWg.Go(func() error {
		h.config.Logger.Debug("API v1: running server at", h.config.Port)
		if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	// listen to SIGTERM on other goroutine
	errWg.Go(func() error {
		<-errCtx.Done()
		return h.Shutdown()
	})

	// wait until all function calls from the goroutines have returned
	// then returns the first non-nil error (if any) from them.
	err := errWg.Wait()
	if errors.Is(err, context.Canceled) || err == nil {
		h.logger.Info("API v2: server quit gracefully")
		return nil
	}
	h.logger.Errorw("API v2: server closed unexpectedly", "error", err)
	return err
}

// Shutdown give the HTTP server the order to shut down, and await for 5 seconds for all connections to gracefully finish
func (h *Handler) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("API v2: server Shutdown: %w", err)
	}

	<-ctx.Done()

	h.logger.Debug("API v2: server exited cleanly")
	return nil
}
