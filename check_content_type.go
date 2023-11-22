package main

import (
	"github.com/gofiber/fiber/v2"
)

func checkContentType(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.
			Status(fiber.StatusNotAcceptable).
			SendString(`"Content-Type: application/json" expected`)
	}

	return c.Next()
}
