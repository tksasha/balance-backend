package main

import (
	"testing"

	. "github.com/tksasha/balance/assert"
)

func TestNewCategory(t *testing.T) {
	category := NewCategory()

	Eq(t, category.Errors.IsEmpty(), true)
}

func TestBuildCategory(t *testing.T) {
	category := BuildCategory("Category Three")

	Eq(t, category.Name, "Category Three")
	Eq(t, category.Errors.IsEmpty(), true)
}

func TestCreateCategory(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when params are valid", func(t *testing.T) {
		category, _ := CreateCategory(db, &categoryParams{"Category One"})

		Eq(t, category.ID, 1)
		Eq(t, category.Name, "Category One")
	})

	t.Run("when `Name` is empty", func(t *testing.T) {
		category, err := CreateCategory(db, &categoryParams{""})

		Eq(t, category.ID, 0)
		Is(t, err, RecordInvalidError)
	})
}

func TestFindCategory(t *testing.T) {
	db := Open()
	defer Close(db)

	t.Run("when Category is exist", func(t *testing.T) {
		expected := Factory(db, "Category").(Category)

		result, _ := FindCategory(db, expected.ID)

		Eq(t, result.ID, expected.ID)
		Eq(t, result.Name, expected.Name)
	})

	t.Run("when Category is not exist", func(t *testing.T) {
		result, err := FindCategory(db, 1335)

		Eq(t, result, nil)
		Is(t, err, RecordNotFoundError)
	})
}
