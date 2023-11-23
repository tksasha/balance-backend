package model_test

import (
	"slices"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/tksasha/balance/model"
)

var (
	subject = new(model.Model)
)

func TestIsValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		assert.Assert(t, subject.IsValid())
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		subject.Errors.Add("attribute", "can't be blank")

		assert.Equal(t, subject.IsValid(), false)
	})
}

func TestIsNotValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		assert.Equal(t, subject.IsNotValid(), false)
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = model.NewErrors()

		subject.Errors.Add("attribute", "is not valid")

		assert.Assert(t, subject.IsNotValid())
	})
}

func TestAddError(t *testing.T) {
	subject.Errors = model.NewErrors()

	subject.AddError("name", "can't be blank")

	messages := subject.Errors.Get("name")

	assert.Assert(t, slices.Contains(messages, "can't be blank"))
}

func TestValidate(t *testing.T) {
	subject := struct {
		model.Model
		Name string `validate:"required"`
	}{}

	subject.Errors = model.NewErrors()

	model.Validate(&subject)

	assert.Assert(t, subject.IsNotValid())

	messages := subject.Errors.Get("name")

	assert.Assert(t, slices.Contains(messages, "required"))
}
