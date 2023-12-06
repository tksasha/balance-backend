package model_test

import (
	"testing"

	"github.com/tksasha/balance/pkg/model"
	"github.com/tksasha/balance/pkg/model/errors"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

var (
	subject = new(model.Model)
)

func TestIsValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = errors.New()

		assert.Assert(t, subject.IsValid())
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = errors.New()

		subject.Errors.Add("attribute", "can't be blank")

		assert.Assert(t, !subject.IsValid())
	})
}

func TestIsNotValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = errors.New()

		assert.Assert(t, subject.IsValid())
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = errors.New()

		subject.Errors.Add("attribute", "is not valid")

		assert.Assert(t, !subject.IsValid())
	})
}

func TestAddError(t *testing.T) {
	subject.Errors = errors.New()

	subject.AddError("name", "can't be blank")

	assert.Assert(t, is.Contains(subject.Errors["errors"]["name"], "can't be blank"))
}

func TestValidate(t *testing.T) {
	subject := struct {
		model.Model
		Name string `validate:"required"`
	}{}

	subject.Errors = errors.New()

	model.Validate(&subject)

	assert.Assert(t, !subject.IsValid())

	assert.Assert(t, is.Contains(subject.Errors["errors"]["name"], "required"))
}
