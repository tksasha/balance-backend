package main_test

import (
	"database/sql"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/tksasha/balance/date"

	. "github.com/tksasha/balance"
)

func Factory(db *sql.DB, model string) any {
	switch model {
	case "Item":
		category := Factory(db, "Category").(Category)

		params := &ItemParams{
			Date:       date.New(2023, 11, 30),
			Formula:    "2+3",
			CategoryID: category.ID,
		}

		item, err := CreateItem(db, params)
		if err == nil {
			return Item(*item)
		}

		panic(err)
	case "Category":
		category, err := CreateCategory(db, &CategoryParams{Name: gofakeit.AppName()})
		if err == nil {
			return Category(*category)
		}

		panic(err)
	default:
		panic("ERROR: UNKNOWN MODEL !!!")
	}
}
