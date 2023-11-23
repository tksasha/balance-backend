package model

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAdd(t *testing.T) {
	errs := NewErrors()

	errs.Add("name", "can't be blank")

	res := errs.Get("name")[0]

	assert.Equal(t, res, "can't be blank")
}

func TestGet(t *testing.T) {
	errs := NewErrors()

	errs.Add("name", "can't be blank")

	res := errs.Get("name")[0]

	assert.Equal(t, res, "can't be blank")
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		errs := NewErrors()

		assert.Assert(t, errs.IsEmpty())
	})

	t.Run("when it is not empty", func(t *testing.T) {
		errs := NewErrors()

		errs.Add("name", "can't be blank")

		assert.Equal(t, errs.IsEmpty(), false)
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		errs := NewErrors()

		assert.Equal(t, errs.IsNotEmpty(), false)
	})

	t.Run("when it is not empty", func(t *testing.T) {
		errs := NewErrors()

		errs.Add("name", "can't be blank")

		assert.Assert(t, errs.IsNotEmpty())
	})
}
