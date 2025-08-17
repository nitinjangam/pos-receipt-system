package repository

import "context"

type SalesRepository struct {
}

func NewSalesRepository(ctx context.Context) *SalesRepository {
	return &SalesRepository{}
}
