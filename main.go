package main

import (
	"log"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	MsgBlank   = "can't be blank"
	MsgInvalid = "is not valid"
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
