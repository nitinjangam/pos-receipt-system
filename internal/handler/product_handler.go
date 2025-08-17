package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
	"github.com/nitinjangam/pos-receipt-system/internal/service"
	"go.uber.org/zap"
)

type ProductHandlerInterface interface {
	GetProducts(c *gin.Context, params v1.GetProductsParams)
	PostProducts(c *gin.Context)
	PutProductsId(c *gin.Context, id int)
	DeleteProductsId(c *gin.Context, id int)
}

type ProductHandler struct {
	productService service.ProductServiceInterface
	logger         *zap.SugaredLogger
}

func NewProductHandler(productService service.ProductServiceInterface, logger *zap.SugaredLogger) ProductHandlerInterface {
	return &ProductHandler{
		productService: productService,
		logger:         logger,
	}
}

func (s *ProductHandler) GetProducts(c *gin.Context, params v1.GetProductsParams) {
	productName := ""
	if params.Name != nil {
		productName = *params.Name
	}
	// get products
	products, err := s.productService.GetProducts(c.Request.Context(), productName)
	if err != nil {
		s.logger.Debugw("Failed to get products", "error", err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{
		"products": products,
	})
}

func (s *ProductHandler) PostProducts(c *gin.Context) {
	// parse products from request body
	var product v1.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		s.logger.Debugw("Failed to bind products", "error", err)
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	// store products
	if err := s.productService.PostProducts(c.Request.Context(), product); err != nil {
		s.logger.Debugw("Failed to post products", "error", err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{"message": "Products created successfully"})
}

func (s *ProductHandler) PutProductsId(c *gin.Context, id int) {
	// parse product from request body
	var product v1.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		s.logger.Debugw("Failed to bind product", "error", err)
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	// update product by id
	updatedProduct, err := s.productService.PutProductsId(c.Request.Context(), product)
	if err != nil {
		s.logger.Debugw("Failed to update product", "error", err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"product": updatedProduct,
	})

}

func (s *ProductHandler) DeleteProductsId(c *gin.Context, id int) {
	// delete product by id
	if err := s.productService.DeleteProductsId(c.Request.Context(), id); err != nil {
		s.logger.Debugw("Failed to delete product", "error", err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(204, gin.H{"message": "Product deleted successfully"})
}
