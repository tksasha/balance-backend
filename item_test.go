package main

import (
	"encoding/json"
	"slices"
	"strings"
	"testing"

	"github.com/tksasha/date"
	"github.com/tksasha/model"
	"gotest.tools/v3/assert"
)

func TestNewItem(t *testing.T) {
	item := NewItem()

	assert.Assert(t, item.Errors.IsEmpty())
}

func TestCalculateItemSum(t *testing.T) {
	t.Run("when `Formula` is empty", func(t *testing.T) {
		item := NewItem()

		item.Formula = ""

		item.Calculate()

		assert.Equal(t, item.Sum, 0.0)
	})

	t.Run("when `Formula` is not valid", func(t *testing.T) {
		item := NewItem()

		item.Formula = "(42.1 + 69.01"

		item.Calculate()

		errs := item.Errors.Get("formula")

		assert.Equal(t, strings.Join(errs, ", "), "is not valid")
	})
}

func TestValidateCategoryID(t *testing.T) {
	t.Run("when it is blank", func(t *testing.T) {
		item := NewItem()

		item.Validate()

		errs := item.Errors.Get("category_id")

		assert.Assert(t, slices.Contains(errs, "can't be blank"))
	})

	t.Run("when it is not blank", func(t *testing.T) {
		item := NewItem()

		item.CategoryID = 1830

		item.Validate()

		errs := item.Errors.Get("category_id")

		assert.Equal(t, strings.Join(errs, ", "), "")
	})
}

func TestValidateFormula(t *testing.T) {
	t.Run("when it is blank", func(t *testing.T) {
		item := NewItem()

		item.Validate()

		errs := item.Errors.Get("formula")

		assert.Assert(t, slices.Contains(errs, "can't be blank"))
	})

	t.Run("when it is not blank", func(t *testing.T) {
		item := NewItem()

		item.Formula = "42.0 + 69.0"

		item.Validate()

		errs := item.Errors.Get("formula")

		assert.Equal(t, strings.Join(errs, ", "), "")
	})
}

func TestValidateDate(t *testing.T) {
	t.Run("when it is blank", func(t *testing.T) {
		item := NewItem()

		item.Validate()

		errs := item.Errors.Get("date")

		assert.Assert(t, slices.Contains(errs, "can't be blank"))
	})

	t.Run("when it is not blank", func(t *testing.T) {
		item := NewItem()

		item.Date = date.Today()

		item.Validate()

		errs := item.Errors.Get("date")

		assert.Equal(t, strings.Join(errs, ", "), "")
	})
}

func TestCreateItem(t *testing.T) {
	db := Open()

	defer Close(db)

	params := new(ItemParams)

	t.Run("when `Date` is blank", func(t *testing.T) {
		item, err := CreateItem(db, params)

		assert.Error(t, err, model.ErrIsNotValid)

		errs := item.Errors.Get("date")

		assert.Assert(t, slices.Contains(errs, "can't be blank"))
	})

	t.Run("when `CategoryID` is zero", func(t *testing.T) {
		item, err := CreateItem(db, params)

		assert.Error(t, err, model.ErrIsNotValid)

		errs := item.Errors.Get("category_id")

		assert.Assert(t, slices.Contains(errs, "can't be blank"))
	})

	t.Run("when `Formula` is not valid", func(t *testing.T) {
		params := &ItemParams{Formula: "(42.1 + 69.01"}

		item, err := CreateItem(db, params)

		assert.Error(t, err, model.ErrIsNotValid)

		errs := item.Errors.Get("formula")

		assert.Assert(t, slices.Contains(errs, "is not valid"))
	})

	t.Run("when `Formula` is empty", func(t *testing.T) {
		item, err := CreateItem(db, params)

		assert.Error(t, err, model.ErrIsNotValid)

		errs := item.Errors.Get("formula")

		assert.Assert(t, slices.Contains(errs, "can't be blank"))
	})

	t.Run("when params are valid", func(t *testing.T) {
		exp := `{"id":1,"date":"2023-11-20","formula":"42.1 + 69.01","sum":111.11,"category_id":2104,"description":"lorem ipsum ..."}`

		params := &ItemParams{
			Date:        date.New(2023, 11, 20),
			Formula:     "42.1 + 69.01",
			CategoryID:  2104,
			Description: "lorem ipsum ...",
		}

		item, err := CreateItem(db, params)

		assert.NilError(t, err)

		res, _ := json.Marshal(item)

		assert.Equal(t, exp, string(res))
	})
}