package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewCategory(t *testing.T) {
	category := NewCategory()

	assert.Assert(t, category.Errors.IsEmpty())
}

func TestBuildCategory(t *testing.T) {
	category := BuildCategory("Category Three")

	assert.Equal(t, category.Name, "Category Three")
	assert.Assert(t, category.Errors.IsEmpty())
}

func TestCreateCategory(t *testing.T) {
	db := Open()

	defer Close(db)

	t.Run("when params are valid", func(t *testing.T) {
		category, _ := CreateCategory(db, &categoryParams{"Category One"})

		assert.Assert(t, category.ID == 1)
		assert.Equal(t, category.Name, "Category One")
	})

	t.Run("when `Name` is empty", func(t *testing.T) {
		category, err := CreateCategory(db, &categoryParams{""})

		assert.Equal(t, category.ID, 0)
		assert.ErrorIs(t, err, RecordInvalidError)
	})
}

func TestFindCategory(t *testing.T) {
	db := Open()

	defer Close(db)

	t.Run("when Category is exist", func(t *testing.T) {
		expected := Factory(db, "Category").(Category)

		result, _ := FindCategory(db, expected.ID)

		assert.Equal(t, result.ID, expected.ID)
		assert.Equal(t, result.Name, expected.Name)
	})

	t.Run("when Category is not exist", func(t *testing.T) {
		result, err := FindCategory(db, 1335)

		assert.Assert(t, result == nil)
		assert.ErrorIs(t, err, RecordNotFoundError)
	})
}
