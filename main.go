package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"koriebruh/cqrs/config"
	"koriebruh/cqrs/internal/command"
	"koriebruh/cqrs/internal/delivery"
	"log"
)

func main() {
	cnf := config.GetConfig()
	db := config.MysqlClient(cnf)
	app := fiber.New()
	validate := validator.New()

	productRepository := command.NewProductRepository()
	productService := command.NewProductService(productRepository, db, validate)
	productHandler := delivery.NewProductHandler(productService)

	app.Get("/", hellobg)
	app.Post("/api/products", productHandler.Create)
	app.Put("/api/products/:id", productHandler.Update)
	app.Delete("/api/products/:id", productHandler.Delete)

	server := fmt.Sprintf("%s:%s", cnf.Server.Host, cnf.Server.Port)
	if err := app.Listen(server); err != nil {
		log.Fatalf("server terminated %v", err)
	}
}

func hellobg(ctx *fiber.Ctx) error {
	return ctx.SendString("woiii")
}
