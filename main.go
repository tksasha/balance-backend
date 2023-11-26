package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	ClientError = errors.New("ClientError")
	ServerError = errors.New("ServerError")
)

func main() {
	app := fiber.New()

	if os.Getenv("GOENV") != "test" {
		app.Use(logger.New())
	}

	app.Use(checkContentType)

	app.Post("/items", CreateItemHandler)
	app.Get("/items/:id", GetItemHandler)

	log.Fatal(app.Listen(":3000"))
}
