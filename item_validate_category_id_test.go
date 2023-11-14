package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/tksasha/validations"
)

func TestValidateCategoryID(t *testing.T) {
	t.Run("when it is blank", func(t *testing.T) {
		item := NewItem()

		item.Validate()

		errors := item.Errors.Get("category_id")

		subject := strings.Join(errors, ", ")

		expected := "can't be blank"

		if !slices.Contains(errors, expected) {
			t.Errorf(validations.M, expected, subject)
		}
	})

	t.Run("when it is not blank", func(t *testing.T) {
		item := NewItem()

		item.CategoryID = 1830

		item.Validate()

		errors := item.Errors.Get("category_id")

		subject := strings.Join(errors, ", ")

		expected := ""

		if expected != subject {
			t.Errorf(validations.M, expected, subject)
		}
	})
}
