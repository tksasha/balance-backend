package main_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/date"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"

	. "github.com/tksasha/balance"
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
	assert.Assert(t, item.CreatedAt.IsZero())
	assert.Assert(t, item.UpdatedAt.IsZero())
}

func TestBuildItem(t *testing.T) {
	item := BuildItem(date.New(2023, 11, 29), "42.1 + 69.01", 23, "lorem ipsum ...")

	assert.Assert(t, item.Errors.IsEmpty())

	assert.Equal(t, item.ID, 0)
	assert.Equal(t, item.Date, date.New(2023, 11, 29))
	assert.Equal(t, item.Formula, "42.1 + 69.01")
	assert.Equal(t, item.Sum, 0.0)
	assert.Equal(t, item.CategoryID, 23)
	assert.Equal(t, item.Description, "lorem ipsum ...")
	assert.Assert(t, item.CreatedAt.IsZero())
	assert.Assert(t, item.UpdatedAt.IsZero())
}

func TestCreateItem(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when params are valid", func(t *testing.T) {
		category := Factory(db, "Category").(Category)

		params := &ItemParams{
			Date:        date.New(2023, 11, 20),
			Formula:     "42.1 + 69.01",
			CategoryID:  category.ID,
			Description: "lorem ipsum ...",
		}

		today := time.Now()

		item, err := CreateItem(db, params)

		assert.NilError(t, err)
		assert.Assert(t, item.Errors.IsEmpty())
		assert.Equal(t, item.ID, 1)
		assert.Equal(t, item.Date, date.New(2023, 11, 20))
		assert.Equal(t, item.Formula, "42.1 + 69.01")
		assert.Equal(t, item.Sum, 111.11)
		assert.Equal(t, item.CategoryID, category.ID)
		assert.Equal(t, item.Description, "lorem ipsum ...")
		assert.Assert(t, item.CreatedAt.After(today))
		assert.Assert(t, item.UpdatedAt.After(today))
	})

	t.Run("when params are not valid", func(t *testing.T) {
		item, err := CreateItem(db, &ItemParams{})

		errs := item.Errors["errors"]

		assert.ErrorIs(t, err, RecordInvalidError)
		assert.Assert(t, is.Contains(errs["date"], "required"))
		assert.Assert(t, is.Contains(errs["formula"], "required"))
		assert.Assert(t, is.Contains(errs["category_id"], "required"))
	})
}

func TestFindItem(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when Item is exist", func(t *testing.T) {
		created := Factory(db, "Item").(Item)

		item, err := FindItem(db, created.ID)

		assert.NilError(t, err)
		assert.Equal(t, item.ID, created.ID)
		assert.Equal(t, item.Date, created.Date)
		assert.Equal(t, item.Formula, created.Formula)
		assert.Equal(t, item.Sum, created.Sum)
		assert.Equal(t, item.CategoryID, created.CategoryID)
		assert.Equal(t, item.Description, created.Description)
		assert.Assert(t, item.CreatedAt.Equal(created.CreatedAt))
		assert.Assert(t, item.UpdatedAt.Equal(created.UpdatedAt))
	})

	t.Run("when item is not exist", func(t *testing.T) {
		item, err := FindItem(db, 1203)

		assert.ErrorIs(t, err, RecordNotFoundError)
		assert.Assert(t, is.Nil(item))
	})
}

func TestUpdateItem(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when params are valid", func(t *testing.T) {
		category := Factory(db, "Category").(Category)
		item := Factory(db, "Item").(Item)

		params := &ItemParams{
			Date:        date.New(2023, 12, 31),
			Formula:     "30.1 + 40.3",
			CategoryID:  category.ID,
			Description: "very random description",
		}

		err := item.Update(db, params)

		assert.NilError(t, err)
		assert.Equal(t, item.Date, date.New(2023, 12, 31))
		assert.Equal(t, item.Formula, "30.1 + 40.3")
		assert.Equal(t, item.Sum, 70.4)
		assert.Equal(t, item.CategoryID, category.ID)
		assert.Equal(t, item.Description, "very random description")
	})

	t.Run("when params are not valid", func(t *testing.T) {
		item := Factory(db, "Item").(Item)

		err := item.Update(db, &ItemParams{})

		errs := item.Errors["errors"]

		assert.ErrorIs(t, err, RecordInvalidError)
		assert.Assert(t, is.Contains(errs["date"], "required"))
		assert.Assert(t, is.Contains(errs["formula"], "required"))
		assert.Assert(t, is.Contains(errs["category_id"], "required"))
	})

	t.Run("when Formula is not valid", func(t *testing.T) {
		item := Factory(db, "Item").(Item)

		err := item.Update(db, &ItemParams{Formula: "(2++"})

		errs := item.Errors["errors"]

		assert.ErrorIs(t, err, RecordInvalidError)
		assert.Assert(t, is.Contains(errs["formula"], "is not valid"))
	})
}
