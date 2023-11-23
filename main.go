package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	ClientError = errors.New("ClientError")
	ServerError = errors.New("ServerError")
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(checkContentType)

	app.Post("/items", CreateItemHandler)

	log.Fatal(app.Listen(":3000"))
}
