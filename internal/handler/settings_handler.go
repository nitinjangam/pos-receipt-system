package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nitinjangam/pos-receipt-system/internal/service"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type SettingsHandlerInterface interface {
	GetSettings(c *gin.Context)
	PutSettings(c *gin.Context)
}

type SettingsHandler struct {
	logger          *zap.SugaredLogger
	tracer          trace.Tracer
	settingsService service.SettingsServiceInterface
}

func NewSettingsHandler(tracer trace.Tracer, logger *zap.SugaredLogger, settingsService service.SettingsServiceInterface) SettingsHandlerInterface {
	return &SettingsHandler{
		logger:          logger,
		tracer:          tracer,
		settingsService: settingsService,
	}
}

func (s *SettingsHandler) GetSettings(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SettingsHandler) PutSettings(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}
