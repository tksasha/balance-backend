package main

import (
	"slices"
	"testing"

	"github.com/tksasha/balance/date"
	"gotest.tools/v3/assert"
)

func TestNewItem(t *testing.T) {
	item := NewItem()

	assert.Assert(t, item.Errors.IsEmpty())
	assert.Equal(t, item.ID, 0)
	assert.Equal(t, item.Date, date.New(1, 1, 1))
	assert.Equal(t, item.Formula, "")
	assert.Equal(t, item.Sum, 0.0)
	assert.Equal(t, item.CategoryID, 0)
	assert.Equal(t, item.Description, "")
}

func TestBuildItem(t *testing.T) {
	item := BuildItem(
		date.New(2023, 11, 29),
		"42.1 + 69.01",
		23,
		"lorem ipsum ...",
	)

	assert.Assert(t, item.Errors.IsEmpty())
	assert.Equal(t, item.ID, 0)
	assert.Equal(t, item.Date, date.New(2023, 11, 29))
	assert.Equal(t, item.Formula, "42.1 + 69.01")
	assert.Equal(t, item.Sum, 0.0)
	assert.Equal(t, item.CategoryID, 23)
	assert.Equal(t, item.Description, "lorem ipsum ...")
}

func TestCreateItem(t *testing.T) {
	db := Open()

	defer Close(db)

	params := new(itemParams)

	t.Run("when `Date` is blank", func(t *testing.T) {
		item, err := CreateItem(db, params)

		assert.ErrorIs(t, err, RecordInvalidError)

		errs := item.Errors["errors"]["date"]

		assert.Assert(t, slices.Contains(errs, "required"))
	})

	t.Run("when `Formula` is empty", func(t *testing.T) {
		item, err := CreateItem(db, params)

		assert.ErrorIs(t, err, RecordInvalidError)

		errs := item.Errors["errors"]["formula"]

		assert.Assert(t, slices.Contains(errs, "required"))
	})

	t.Run("when `Formula` is not valid", func(t *testing.T) {
		params := &itemParams{Formula: "(42.1 + 69.01"}

		item, err := CreateItem(db, params)

		assert.ErrorIs(t, err, RecordInvalidError)

		errs := item.Errors["errors"]["formula"]

		assert.Assert(t, slices.Contains(errs, "is not valid"))
	})

	t.Run("when `CategoryID` is zero", func(t *testing.T) {
		item, err := CreateItem(db, params)

		assert.ErrorIs(t, err, RecordInvalidError)

		errs := item.Errors["errors"]["category_id"]

		assert.Assert(t, slices.Contains(errs, "required"))
	})

	t.Run("when params are valid", func(t *testing.T) {
		category, _ := CreateCategory(db, &categoryParams{"Category Fourth"})

		params := &itemParams{
			Date:        date.New(2023, 11, 20),
			Formula:     "42.1 + 69.01",
			CategoryID:  category.ID,
			Description: "lorem ipsum ...",
		}

		item, err := CreateItem(db, params)

		assert.NilError(t, err)

		assert.Assert(t, item.Errors.IsEmpty())
		assert.Equal(t, item.ID, 1)
		assert.Equal(t, item.Date, date.New(2023, 11, 20))
		assert.Equal(t, item.Formula, "42.1 + 69.01")
		assert.Equal(t, item.Sum, 111.11)
		assert.Equal(t, item.CategoryID, category.ID)
		assert.Equal(t, item.Description, "lorem ipsum ...")
	})
}
