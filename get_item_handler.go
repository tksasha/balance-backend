package main

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetItemHandler(c *fiber.Ctx) error {
	db := Open()

	defer Close(db)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.
			Status(fiber.StatusNotFound).
			JSON("")
	}

	item, err := GetItemQuery(db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.
				Status(fiber.StatusNotFound).
				JSON("")
		} else {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(err)
		}
	}

	return c.
		Status(fiber.StatusOK).
		JSON(item)
}
