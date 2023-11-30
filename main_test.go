package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/tksasha/balance/date"
)

func TestMain(m *testing.M) {
	if os.Getenv("GOENV") != "test" {
		panic("What are You doing?!")
	}

	code := func(m *testing.M) int {
		db := Open()

		truncate(db)

		defer func() {
			truncate(db)

			Close(db)
		}()

		return m.Run()
	}(m)

	os.Exit(code)
}

func Factory(db *sql.DB, model string) any {
	switch model {
	case "item", "Item":
		category := Factory(db, "Category").(Category)

		params := &itemParams{
			Date:       date.New(2023, 11, 30),
			Formula:    "2+3",
			CategoryID: category.ID,
		}

		item, err := CreateItem(db, params)

		if err == nil {
			return Item(*item)
		}

		panic(err)
	case "category", "Category":
		category, err := CreateCategory(db, &categoryParams{Name: gofakeit.AppName()})
		if err == nil {
			return Category(*category)
		}

		panic(err)
	default:
		panic("unknown model")
	}
}

func truncate(db *sql.DB) {
	for _, table := range []string{"categories", "cashes", "items"} {
		db.Exec(fmt.Sprintf("DELETE FROM %s", table))
	}
}
