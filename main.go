package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(checkContentType)

	app.Post("/items", CreateItemHandler)

	log.Fatal(app.Listen(":3000"))
}
