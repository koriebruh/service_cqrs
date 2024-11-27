package command

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"koriebruh/cqrs/dto"
	"koriebruh/cqrs/internal/domain"
	"koriebruh/cqrs/pkg/helper"
	"net/http"
)

type ProductService interface {
	Create(ctx context.Context, create dto.ProductCreate) (dto.WebRes, error)
	Update(ctx context.Context, update dto.ProductUpdate) (dto.WebRes, error)
	Delete(ctx context.Context, id int) (dto.WebRes, error)
}

type ProductServiceImpl struct {
	*ProductRepository
	*gorm.DB
	*validator.Validate
}

func NewProductService(productRepository *ProductRepository, DB *gorm.DB, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{ProductRepository: productRepository, DB: DB, Validate: validate}
}

func (service ProductServiceImpl) Create(ctx context.Context, create dto.ProductCreate) (dto.WebRes, error) {
	if err := service.Validate.Struct(create); err != nil {
		return helper.ErrorResponseMsg(helper.ErrBadRequest, err), nil
	}

	product := new(domain.Product)

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		product = &domain.Product{
			Name:        create.Name,
			Description: create.Description,
			Price:       create.Price,
			Stock:       create.Stock,
		}

		err := service.ProductRepository.Create(ctx, tx, product)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return helper.ErrorResponseMsg(helper.ErrBadRequest, err), nil
	}

	return helper.SuccessRes(http.StatusCreated, "OK", "success created new product"), nil

}

func (service ProductServiceImpl) Update(ctx context.Context, update dto.ProductUpdate) (dto.WebRes, error) {
	//TODO implement me
	panic("implement me")
}

func (service ProductServiceImpl) Delete(ctx context.Context, id int) (dto.WebRes, error) {
	//TODO implement me
	panic("implement me")
}
