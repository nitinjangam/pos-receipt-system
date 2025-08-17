package service

import (
	"context"
	"errors"

	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
	"github.com/nitinjangam/pos-receipt-system/internal/repository"
	"go.uber.org/zap"
)

type ProductServiceInterface interface {
	GetProducts(ctx context.Context, productName string) ([]v1.Product, error)
	PostProducts(ctx context.Context, products v1.Product) error
	PutProductsId(ctx context.Context, product v1.Product) (v1.Product, error)
	DeleteProductsId(ctx context.Context, id int) error
}

type ProductService struct {
	productRepo *repository.ProductRepository
	logger      *zap.SugaredLogger
}

func NewProductService(productRepository *repository.ProductRepository, logger *zap.SugaredLogger) *ProductService {
	return &ProductService{
		productRepo: productRepository,
		logger:      logger,
	}
}

func (s *ProductService) GetProducts(ctx context.Context, prodctName string) ([]v1.Product, error) {
	if prodctName != "" {
		// Get product by name from the repository
		product, err := s.productRepo.GetProductByName(ctx, prodctName)
		if err != nil {
			s.logger.Debugw("Failed to get product by name", "error", err, "product_name", prodctName)
			return nil, err
		}
		if product == nil {
			s.logger.Debugw("Product not found", "product_name", prodctName)
			return nil, errors.New("product not found")
		}
		// Return a slice with the single product
		return []v1.Product{*product}, nil
	}
	// Get all products from the repository
	products, err := s.productRepo.GetAllProducts(ctx)
	if err != nil {
		s.logger.Debugw("Failed to get products", "error", err)
		return nil, err
	}

	if len(products) == 0 {
		s.logger.Debugw("No products found")
		return nil, nil // or return an empty slice if preferred
	}

	return products, nil
}

func (s *ProductService) PostProducts(ctx context.Context, product v1.Product) error {
	// Check if the product already exists
	existingProduct, err := s.productRepo.GetProductByName(ctx, *product.Name)
	if err != nil {
		s.logger.Debugw("Failed to get product by name", "error", err, "product_name", *product.Name)
		return err
	}
	if existingProduct != nil {
		s.logger.Debugw("Product already exists", "product_name", *product.Name)
	}
	if err := s.productRepo.CreateProduct(ctx, product); err != nil {
		s.logger.Debugw("Failed to create product", "error", err, "product", product)
		return err
	}
	return nil
}

func (s *ProductService) PutProductsId(ctx context.Context, product v1.Product) (v1.Product, error) {
	// Check if the product exists
	existingProduct, err := s.productRepo.GetProductByID(ctx, *product.Id)
	if err != nil {
		s.logger.Debugw("Failed to get product by ID", "error", err, "product_id", product.Id)
		return v1.Product{}, err
	}
	if existingProduct == nil {
		s.logger.Debugw("Product not found", "product_id", product.Id)
		return v1.Product{}, errors.New("product not found")
	}

	// Update the product in the repository
	if err := s.productRepo.UpdateProduct(ctx, product); err != nil {
		s.logger.Debugw("Failed to update product", "error", err, "product", product)
		return v1.Product{}, err
	}

	return product, nil
}

func (s *ProductService) DeleteProductsId(ctx context.Context, id int) error {
	// Check if the product exists
	existingProduct, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		s.logger.Debugw("Failed to get product by ID", "error", err, "product_id", id)
		return err
	}
	if existingProduct == nil {
		s.logger.Debugw("Product not found", "product_id", id)
		return errors.New("product not found")
	}

	// Delete the product from the repository
	if err := s.productRepo.DeleteProduct(ctx, id); err != nil {
		s.logger.Debugw("Failed to delete product", "error", err, "product_id", id)
		return err
	}

	return nil
}
