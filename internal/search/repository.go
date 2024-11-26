package search

import (
	"context"
	"koriebruh/cqrs/internal/domain"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r ProductRepository) FindById(ctx context.Context, id int) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r ProductRepository) FindAll(ctx context.Context, page int) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
