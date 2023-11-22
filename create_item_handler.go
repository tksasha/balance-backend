package main

import (
	"github.com/gofiber/fiber/v2"
)

func CreateItemHandler(c *fiber.Ctx) error {
	db := Open()

	defer Close(db)

	params := new(ItemParams)

	if err := c.BodyParser(params); err != nil {
		return c.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.ErrUnprocessableEntity)
	}

	item, err := CreateItem(db, params)
	if err == nil {
		return c.
			Status(fiber.StatusCreated).
			JSON(item)
	} else {
		return c.
			Status(fiber.StatusUnprocessableEntity).
			JSON(item.Errors)
	}
}
