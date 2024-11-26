package command

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/cqrs/internal/domain"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r ProductRepository) Create(ctx context.Context, db *gorm.DB, product *domain.Product) error {
	create := db.WithContext(ctx).Create(product)
	if create.Error != nil {
		return fmt.Errorf("error create new product")
	}

	return nil
}

func (r ProductRepository) Update(ctx context.Context, db *gorm.DB, id int, product *domain.Product) error {
	var existingProduct domain.Product
	if err := db.WithContext(ctx).First(&existingProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("product with id %d not found", id)
		}
		return fmt.Errorf("error fetching product with id %d: %w", id, err)
	}

	if err := db.WithContext(ctx).Model(&existingProduct).Updates(product).Error; err != nil {
		return fmt.Errorf("error updating product with id %d: %w", id, err)
	}

	return nil
}

func (r ProductRepository) Delete(ctx context.Context, db *gorm.DB, id int) error {
	var existingProduct domain.Product
	if err := db.WithContext(ctx).First(&existingProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("product with id %d not found", id)
		}
		return fmt.Errorf("error finding product with id %d: %w", id, err)
	}

	if err := db.WithContext(ctx).Where("id=?", id).Delete(domain.Product{}).Error; err != nil {
		return fmt.Errorf("error delete product with id %d: %w", id, err)
	}

	return nil
}
