package model_test

import (
	"testing"

	. "github.com/tksasha/balance/assert"
	"github.com/tksasha/balance/model"
)

var (
	subject = new(model.Model)
)

func TestIsValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		Eq(t, subject.IsValid(), true)
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		subject.Errors.Add("attribute", "can't be blank")

		Eq(t, subject.IsValid(), false)
	})
}

func TestIsNotValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		Eq(t, subject.IsValid(), true)
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		subject.Errors.Add("attribute", "is not valid")

		Eq(t, subject.IsValid(), false)
	})
}

func TestAddError(t *testing.T) {
	subject.Errors = model.NewErrors()

	subject.AddError("name", "can't be blank")

	In(t, subject.Errors["errors"]["name"], "can't be blank")
}

func TestValidate(t *testing.T) {
	subject := struct {
		model.Model
		Name string `validate:"required"`
	}{}

	subject.Errors = model.NewErrors()

	model.Validate(&subject)

	Eq(t, subject.IsValid(), false)
	In(t, subject.Errors["errors"]["name"], "required")
}
