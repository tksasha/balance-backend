package main

import (
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/tksasha/validations"
)

func TestValidateDateWhenItIsBlank(t *testing.T) {
	item := NewItem()

	item.Validate()

	errors := item.Errors.Get("date")

	subject := strings.Join(errors, ", ")

	expected := "can't be blank"

	if !slices.Contains(errors, "can't be blank") {
		t.Errorf(validations.MESSAGE, expected, subject)
	}
}

func TestValidateDateWhenItIsNotBlank(t *testing.T) {
	item := NewItem()

	item.Date = time.Now()

	item.Validate()

	errors := item.Errors.Get("date")

	subject := strings.Join(errors, ", ")

	expected := ""

	if len(errors) > 0 {
		t.Errorf(validations.MESSAGE, expected, subject)
	}
}
