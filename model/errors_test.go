package model_test

import (
	"slices"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/tksasha/balance/model"
)

func TestAdd(t *testing.T) {
	t.Run("when attribute in lowercase", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("name", "required")

		messages := subject.Get("name")

		assert.Assert(t, slices.Contains(messages, "required"))
	})

	t.Run("when attribute in uppercase", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("Name", "required")

		messages := subject.Get("name")

		assert.Assert(t, slices.Contains(messages, "required"))
	})
}

func TestGet(t *testing.T) {
	subject := model.NewErrors()

	subject.Add("name", "can't be blank")

	messages := subject.Get("name")

	assert.Assert(t, slices.Contains(messages, "can't be blank"))
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		subject := model.NewErrors()

		assert.Assert(t, subject.IsEmpty())
	})

	t.Run("when it is not empty", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("name", "can't be blank")

		assert.Equal(t, subject.IsEmpty(), false)
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		subject := model.NewErrors()

		assert.Equal(t, subject.IsNotEmpty(), false)
	})

	t.Run("when it is not empty", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("name", "can't be blank")

		assert.Assert(t, subject.IsNotEmpty())
	})
}
