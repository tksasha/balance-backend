package errors_test

import (
	"testing"

	"github.com/tksasha/balance/pkg/model/errors"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNew(t *testing.T) {
	t.Run("when no one argument is provided", func(t *testing.T) {
		subject := errors.New()

		assert.Assert(t, subject.IsEmpty())
	})

	t.Run("when only one argument is provided", func(t *testing.T) {
		subject := errors.New("name")

		assert.Assert(t, subject.IsEmpty())
	})

	t.Run("when two arguments are provided", func(t *testing.T) {
		subject := errors.New("name", "can't be blank")

		assert.Assert(t, is.Contains(subject["errors"]["name"], "can't be blank"))
	})
}

func TestAdd(t *testing.T) {
	t.Run("when attribute in lowercase", func(t *testing.T) {
		subject := errors.New()

		subject.Add("name", "required")

		assert.Assert(t, is.Contains(subject["errors"]["name"], "required"))
	})

	t.Run("when attribute in uppercase", func(t *testing.T) {
		subject := errors.New()

		subject.Add("Name", "required")

		assert.Assert(t, is.Contains(subject["errors"]["name"], "required"))
	})
}

func TestGet(t *testing.T) {
	subject := errors.New()

	subject.Add("name", "can't be blank")

	assert.Assert(t, is.Contains(subject["errors"]["name"], "can't be blank"))
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		subject := errors.New()

		assert.Assert(t, subject.IsEmpty())
	})

	t.Run("when it is not empty", func(t *testing.T) {
		subject := errors.New()

		subject.Add("name", "can't be blank")

		assert.Assert(t, !subject.IsEmpty())
	})
}
