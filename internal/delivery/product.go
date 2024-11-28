package delivery

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/cqrs/dto"
	"koriebruh/cqrs/internal/command"
	"koriebruh/cqrs/pkg/helper"
	"net/http"
	"strconv"
)

type ProductHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type ProductHandlerImpl struct {
	command.ProductService
}

func NewProductHandler(productService command.ProductService) *ProductHandlerImpl {
	return &ProductHandlerImpl{ProductService: productService}
}

func (handler ProductHandlerImpl) Create(ctx *fiber.Ctx) error {
	var req dto.ProductCreate
	if err := ctx.BodyParser(&req); err != nil {
		err = helper.ErrInternalServerErr
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	created, err := handler.ProductService.Create(ctx.Context(), req)
	if err != nil {
		err = helper.ErrInternalServerErr
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(created)
}

func (handler ProductHandlerImpl) Update(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		panic(err)
	}

	var req dto.ProductUpdate
	if err := ctx.BodyParser(&req); err != nil {
		err = helper.ErrInternalServerErr
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	updated, err := handler.ProductService.Update(ctx.Context(), id, req)
	if err != nil {
		err = helper.ErrInternalServerErr
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(updated)
}

func (handler ProductHandlerImpl) Delete(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		panic(err)
	}

	res, err := handler.ProductService.Delete(ctx.Context(), id)
	if err != nil {
		err = helper.ErrInternalServerErr
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
