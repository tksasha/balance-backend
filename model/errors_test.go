package model_test

import (
	"slices"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/tksasha/balance/model"
)

func TestNewErrors(t *testing.T) {
	t.Run("when no one argument is provided", func(t *testing.T) {
		subject := model.NewErrors()

		assert.Assert(t, subject.IsEmpty())
	})

	t.Run("when only one argument is provided", func(t *testing.T) {
		subject := model.NewErrors("name")

		assert.Assert(t, subject.IsEmpty())
	})

	t.Run("when two arguments are provided", func(t *testing.T) {
		subject := model.NewErrors("name", "can't be blank")

		assert.Assert(t, slices.Contains(subject["errors"]["name"], "can't be blank"))
	})
}

func TestAdd(t *testing.T) {
	t.Run("when attribute in lowercase", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("name", "required")

		assert.Assert(t, slices.Contains(subject["errors"]["name"], "required"))
	})

	t.Run("when attribute in uppercase", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("Name", "required")

		assert.Assert(t, slices.Contains(subject["errors"]["name"], "required"))
	})
}

func TestGet(t *testing.T) {
	subject := model.NewErrors()

	subject.Add("name", "can't be blank")

	assert.Assert(t, slices.Contains(subject["errors"]["name"], "can't be blank"))
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		subject := model.NewErrors()

		assert.Assert(t, subject.IsEmpty())
	})

	t.Run("when it is not empty", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("name", "can't be blank")

		assert.Assert(t, !subject.IsEmpty())
	})
}
