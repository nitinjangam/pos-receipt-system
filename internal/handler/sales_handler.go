package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nitinjangam/pos-receipt-system/internal/service"
	"go.uber.org/zap"
)

// SalesHandlerInterface defines the methods for the sales service.
type SalesHandlerInterface interface {
	GetSales(c *gin.Context)
	PostSales(c *gin.Context)
	DeleteSalesId(c *gin.Context, id int)
	PutSalesId(c *gin.Context, id int)
	GetSalesIdReceipt(c *gin.Context, id int)
}

type SalesHandler struct {
	ctx          context.Context
	logger       *zap.SugaredLogger
	salesService service.SalesServiceInterface
}

func NewSalesHandler(ctx context.Context, logger *zap.SugaredLogger, salesService service.SalesServiceInterface) SalesHandlerInterface {
	return &SalesHandler{
		ctx:          ctx,
		logger:       logger,
		salesService: salesService,
	}
}

func (s *SalesHandler) GetSales(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesHandler) PostSales(c *gin.Context) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesHandler) DeleteSalesId(c *gin.Context, id int) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesHandler) PutSalesId(c *gin.Context, id int) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func (s *SalesHandler) GetSalesIdReceipt(c *gin.Context, id int) {
	// return status not implemented
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}
