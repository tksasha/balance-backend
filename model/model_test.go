package model_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/tksasha/balance/model"
)

func TestIsValid(t *testing.T) {
	subject := new(model.Model)

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
	subject := new(model.Model)

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
