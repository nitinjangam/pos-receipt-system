package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nitinjangam/pos-receipt-system/internal/repository"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type SettingsServiceInterface interface {
	GetSettings(c *gin.Context)
	PutSettings(c *gin.Context)
}

type SettingsService struct {
	logger       *zap.SugaredLogger
	tracer       trace.Tracer
	settingsRepo *repository.SettingsRepository
}

func NewSettingsService(tracer trace.Tracer, logger *zap.SugaredLogger, settingsRepository *repository.SettingsRepository) *SettingsService {
	return &SettingsService{
		logger:       logger,
		tracer:       tracer,
		settingsRepo: settingsRepository,
	}
}

func (s *SettingsService) GetSettings(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SettingsService) PutSettings(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}
