package main

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/tksasha/balance/date"

	. "github.com/tksasha/balance/assert"
)

func TestNewItem(t *testing.T) {
	item := NewItem()

	Eq(t, item.Errors.IsEmpty(), true)
	Eq(t, item.ID, 0)
	Eq(t, item.Date, date.New(1, 1, 1))
	Eq(t, item.Formula, "")
	Eq(t, item.Sum, 0.0)
	Eq(t, item.CategoryID, 0)
	Eq(t, item.Description, "")
}

func TestBuildItem(t *testing.T) {
	item := BuildItem(
		date.New(2023, 11, 29),
		"42.1 + 69.01",
		23,
		"lorem ipsum ...",
	)

	Eq(t, item.Errors.IsEmpty(), true)
	Eq(t, item.ID, 0)
	Eq(t, item.Date, date.New(2023, 11, 29))
	Eq(t, item.Formula, "42.1 + 69.01")
	Eq(t, item.Sum, 0.0)
	Eq(t, item.CategoryID, 23)
	Eq(t, item.Description, "lorem ipsum ...")
}

func TestCreateItem(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when params are valid", func(t *testing.T) {
		category := Factory(db, "category").(Category)

		params := &itemParams{
			Date:        date.New(2023, 11, 20),
			Formula:     "42.1 + 69.01",
			CategoryID:  category.ID,
			Description: "lorem ipsum ...",
		}

		today := time.Now()

		item, err := CreateItem(db, params)

		Eq(t, err, nil)
		Eq(t, item.Errors.IsEmpty(), true)
		Eq(t, item.ID, 1)
		Eq(t, item.Date, date.New(2023, 11, 20))
		Eq(t, item.Formula, "42.1 + 69.01")
		Eq(t, item.Sum, 111.11)
		Eq(t, item.CategoryID, category.ID)
		Eq(t, item.Description, "lorem ipsum ...")
		Eq(t, item.CreatedAt.After(today), true)
		Eq(t, item.UpdatedAt.After(today), true)
	})
}

func TestFindItem(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when Item is exist", func(t *testing.T) {
		created := Factory(db, "item").(Item)

		item, err := FindItem(db, created.ID)

		Eq(t, err, nil)
		Eq(t, item.ID, created.ID)
	})

	t.Run("when item is not exist", func(t *testing.T) {
		item, err := FindItem(db, 1203)

		Is(t, err, RecordNotFoundError)
		Eq(t, item, nil)
	})
}

func TestUpdateItem(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when params are valid", func(t *testing.T) {
		category := Factory(db, "Category").(Category)
		item := Factory(db, "Item").(Item)
		description := gofakeit.Sentence(10)

		params := &itemParams{
			Date:        date.New(2023, 12, 31),
			Formula:     "30.1 + 40.3",
			CategoryID:  category.ID,
			Description: description,
		}

		err := item.Update(db, params)

		Eq(t, err, nil)
		Eq(t, item.Date, date.New(2023, 12, 31))
		Eq(t, item.Formula, "30.1 + 40.3")
		Eq(t, item.Sum, 70.4)
		Eq(t, item.CategoryID, category.ID)
		Eq(t, item.Description, description)
		Eq(t, item.Description, description)
	})

	t.Run("when params are not valid", func(t *testing.T) {
		item := Factory(db, "Item").(Item)

		err := item.Update(db, &itemParams{})

		errs := item.Errors["errors"]

		Is(t, err, RecordInvalidError)
		In(t, errs["date"], "required")
		In(t, errs["formula"], "required")
		In(t, errs["category_id"], "required")
	})

	t.Run("when Formula is not valid", func(t *testing.T) {
		item := Factory(db, "item").(Item)

		err := item.Update(db, &itemParams{Formula: "(2++"})

		Is(t, err, RecordInvalidError)
		In(t, item.Errors["errors"]["formula"], "is not valid")
	})
}
