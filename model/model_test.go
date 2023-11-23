package model

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIsValid(t *testing.T) {
	model := new(Model)

	t.Run("when it is valid", func(t *testing.T) {
		model.Errors = NewErrors()

		assert.Assert(t, model.IsValid())
	})

	t.Run("when it is not valid", func(t *testing.T) {
		model.Errors = NewErrors()

		model.Errors.Add("attribute", "can't be blank")

		assert.Equal(t, model.IsValid(), false)
	})
}

func TestIsNotValid(t *testing.T) {
	model := new(Model)

	t.Run("when it is valid", func(t *testing.T) {
		model.Errors = NewErrors()

		assert.Equal(t, model.IsNotValid(), false)
	})

	t.Run("when it is not valid", func(t *testing.T) {
		model.Errors = NewErrors()

		model.Errors.Add("attribute", "is not valid")

		assert.Assert(t, model.IsNotValid())
	})
}
