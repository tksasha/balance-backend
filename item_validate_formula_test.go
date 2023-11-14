package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/tksasha/validations"
)

func TestValidateFormula(t *testing.T) {
	t.Run("when it is blank", func(t *testing.T) {
		item := NewItem()

		item.Validate()

		errors := item.Errors.Get("formula")

		subject := strings.Join(errors, ", ")

		expected := "can't be blank"

		if !slices.Contains(errors, expected) {
			t.Errorf(validations.M, expected, subject)
		}
	})

	t.Run("when it is not blank", func(t *testing.T) {
		item := NewItem()

		item.Formula = "42.0 + 69.0"

		item.Validate()

		errors := item.Errors.Get("formula")

		subject := strings.Join(errors, ", ")

		expected := ""

		if len(errors) > 0 {
			t.Errorf(validations.M, expected, subject)
		}
	})
}
