package domain

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	ID          int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"not null; size:150"`
	Description string  `gorm:"size:250"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

type ProductCmdRepository interface {
	Create(ctx context.Context, db *gorm.DB, product *Product) error
	Update(ctx context.Context, db *gorm.DB, id int, product *Product) error
	Delete(ctx context.Context, db *gorm.DB, id int) error
}

type ProductQueryRepository interface {
	FindById(ctx context.Context, id int) (*Product, error)
	FindAll(ctx context.Context, page int) ([]Product, error)
}
