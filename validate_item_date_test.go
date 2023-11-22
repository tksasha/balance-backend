package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/tksasha/date"
)

func TestValidateDate(t *testing.T) {
	t.Run("when it is blank", func(t *testing.T) {
		item := NewItem()

		item.Validate()

		errors := item.Errors.Get("date")

		subject := strings.Join(errors, ", ")

		expected := "can't be blank"

		if !slices.Contains(errors, expected) {
			t.Errorf(M, expected, subject)
		}
	})

	t.Run("when it is not blank", func(t *testing.T) {
		item := NewItem()

		item.Date = date.Today()

		item.Validate()

		errors := item.Errors.Get("date")

		subject := strings.Join(errors, ", ")

		expected := ""

		if len(errors) > 0 {
			t.Errorf(M, expected, subject)
		}
	})
}
