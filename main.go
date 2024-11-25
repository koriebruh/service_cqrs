package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"koriebruh/cqrs/config"
	"log"
)

func main() {
	cnf := config.GetConfig()
	app := fiber.New()

	app.Get("/", hellobg)

	server := fmt.Sprintf("%s:%s", cnf.Server.Host, cnf.Server.Port)
	if err := app.Listen(server); err != nil {
		log.Fatalf("server terminated %v", err)
	}
}

func hellobg(ctx *fiber.Ctx) error {
	return ctx.SendString("woiii")
}
