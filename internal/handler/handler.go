package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
)

type HandlerInterface interface {
	PostAuthLogin(c *gin.Context)
	PostAuthRegister(c *gin.Context)
	GetProducts(c *gin.Context, params v1.GetProductsParams)
	PostProducts(c *gin.Context)
	PutProductsId(c *gin.Context, id int)
	DeleteProductsId(c *gin.Context, id int)
	GetSales(c *gin.Context)
	PostSales(c *gin.Context)
	DeleteSalesId(c *gin.Context, id int)
	PutSalesId(c *gin.Context, id int)
	GetSalesIdReceipt(c *gin.Context, id int)
	GetSettings(c *gin.Context)
	PutSettings(c *gin.Context)
}

type Handler struct {
	AuthHandler     AuthHandlerInterface
	ProductHandler  ProductHandlerInterface
	SalesHandler    SalesHandlerInterface
	SettingsHandler SettingsHandlerInterface
}

func NewHandler(AuthHandler AuthHandlerInterface,
	ProductHandler ProductHandlerInterface,
	SalesHandler SalesHandlerInterface,
	SettingsHandler SettingsHandlerInterface) HandlerInterface {
	return &Handler{
		AuthHandler:     AuthHandler,
		ProductHandler:  ProductHandler,
		SalesHandler:    SalesHandler,
		SettingsHandler: SettingsHandler,
	}
}

// PostAuthLogin handles user login.
func (s *Handler) PostAuthLogin(c *gin.Context) {
	s.AuthHandler.PostAuthLogin(c)
}

// PostAuthRegister handles user registration.
func (s *Handler) PostAuthRegister(c *gin.Context) {
	s.AuthHandler.PostAuthRegister(c)
}

// GetProducts retrieves all products.
func (s *Handler) GetProducts(c *gin.Context, params v1.GetProductsParams) {
	s.ProductHandler.GetProducts(c, params)
}

// PostProducts creates a new product.
func (s *Handler) PostProducts(c *gin.Context) {
	s.ProductHandler.PostProducts(c)
}

// PutProductsId updates a product by ID.
func (s *Handler) PutProductsId(c *gin.Context, id int) {
	s.ProductHandler.PutProductsId(c, id)
}

// DeleteProductsId deletes a product by ID.
func (s *Handler) DeleteProductsId(c *gin.Context, id int) {
	s.ProductHandler.DeleteProductsId(c, id)
}

// GetSales retrieves all sales.
func (s *Handler) GetSales(c *gin.Context) {
	s.SalesHandler.GetSales(c)
}

// PostSales creates a new sale.
func (s *Handler) PostSales(c *gin.Context) {
	s.SalesHandler.PostSales(c)
}

// DeleteSalesId deletes a sale by ID.
func (s *Handler) DeleteSalesId(c *gin.Context, id int) {
	s.SalesHandler.DeleteSalesId(c, id)
}

// PutSalesId updates a sale by ID.
func (s *Handler) PutSalesId(c *gin.Context, id int) {
	s.SalesHandler.PutSalesId(c, id)
}

// GetSalesIdReceipt retrieves a sale receipt by ID.
func (s *Handler) GetSalesIdReceipt(c *gin.Context, id int) {
	s.SalesHandler.GetSalesIdReceipt(c, id)
}

// GetSettings retrieves the settings.
func (s *Handler) GetSettings(c *gin.Context) {
	s.SettingsHandler.GetSettings(c)
}

// PutSettings updates the settings.
func (s *Handler) PutSettings(c *gin.Context) {
	s.SettingsHandler.PutSettings(c)
}
