package repository

import (
	"context"
	"database/sql"

	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
)

// ProductRepositoryInterface defines the methods for the product repository.
type ProductRepositoryInterface interface {
	GetProductByName(ctx context.Context, name string) (*v1.Product, error)
	GetAllProducts(ctx context.Context) ([]v1.Product, error)
	GetProductByID(ctx context.Context, id int) (*v1.Product, error)
	CreateProduct(ctx context.Context, product v1.Product) error
	UpdateProduct(ctx context.Context, product v1.Product) error
	DeleteProduct(ctx context.Context, id int) error
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]v1.Product, error) {
	var products []v1.Product

	query := "SELECT id, name, price, description, cgst_rate, sgst_rate FROM products"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product v1.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.CgstRate, &product.SgstRate); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByName(ctx context.Context, name string) (*v1.Product, error) {
	var product v1.Product

	query := "SELECT id, name, price, description, cgst_rate, sgst_rate FROM products WHERE name = ?"
	err := r.db.QueryRowContext(ctx, query, name).Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.CgstRate, &product.SgstRate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Product not found
		}
		return nil, err // Other error
	}

	return &product, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*v1.Product, error) {
	var product v1.Product

	query := "SELECT id, name, price, description, cgst_rate, sgst_rate FROM products WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.CgstRate, &product.SgstRate)
	if err != nil {
		if err == sql.ErrNoRows {
			return &product, nil // Product not found
		}
		return &product, err // Other error
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product v1.Product) error {
	query := "INSERT INTO products (name, description, price, cgst_rate, sgst_rate) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.CgstRate, product.SgstRate)
	if err != nil {
		return err // Return error if insertion fails
	}
	return nil // Return nil if insertion is successful
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, product v1.Product) error {
	query := "UPDATE products SET name = ?, price = ?, description = ?, sgst_rate = ?, cgst_rate = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.Description, product.SgstRate, product.CgstRate, product.Id)
	if err != nil {
		return err // Return error if update fails
	}
	return nil // Return nil if update is successful
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err // Return error if deletion fails
	}
	return nil // Return nil if deletion is successful
}
