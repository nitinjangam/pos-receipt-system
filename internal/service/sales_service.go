package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nitinjangam/pos-receipt-system/internal/repository"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// SalesServiceInterface defines the methods for the sales service.
type SalesServiceInterface interface {
	GetSales(c *gin.Context)
	PostSales(c *gin.Context)
	DeleteSalesId(c *gin.Context, id int)
	PutSalesId(c *gin.Context, id int)
	GetSalesIdReceipt(c *gin.Context, id int)
}

type SalesService struct {
	logger          *zap.SugaredLogger
	tracer          trace.Tracer
	salesRepository *repository.SalesRepository
}

func NewSalesService(tracer trace.Tracer, logger *zap.SugaredLogger, salesRepository *repository.SalesRepository) *SalesService {
	return &SalesService{
		logger:          logger,
		tracer:          tracer,
		salesRepository: salesRepository,
	}
}

func (s *SalesService) GetSales(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesService) PostSales(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesService) DeleteSalesId(c *gin.Context, id int) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesService) PutSalesId(c *gin.Context, id int) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesService) GetSalesIdReceipt(c *gin.Context, id int) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}
