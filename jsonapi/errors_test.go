package jsonapi_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/tksasha/balance/jsonapi"
)

func TestAdd(t *testing.T) {
	t.Run("when `source` is a `pointer`", func(t *testing.T) {
		subject := jsonapi.NewErrors("pointer", "/data/attributes/name", "required")[0]

		assert.Equal(t, subject.Title, "required")
		assert.Equal(t, subject.Detail, "Can't be blank")
		assert.Equal(t, subject.Source.Pointer, "/data/attributes/name")
	})

	t.Run("when `source` is a `parameter`", func(t *testing.T) {
		subject := jsonapi.NewErrors("parameter", "id", "invalid", "Only integers allowed")[0]

		assert.Equal(t, subject.Title, "invalid")
		assert.Equal(t, subject.Detail, "Only integers allowed")
		assert.Equal(t, subject.Source.Parameter, "id")
	})

	t.Run("when `source` is not specified", func(t *testing.T) {
		subject := jsonapi.NewErrors("server error")[0]

		assert.Equal(t, subject.Title, "server error")
		assert.Equal(t, subject.Detail, "Internal Server Error")
	})

	t.Run("when `title` is `required`", func(t *testing.T) {
		errs := jsonapi.NewErrors("pointer", "/data/attributes/name", "required")

		subject := errs[0]

		assert.Equal(t, subject.Title, "required")
		assert.Equal(t, subject.Detail, "Can't be blank")
	})

	t.Run("when `title` is `invalid`", func(t *testing.T) {
		errs := jsonapi.NewErrors("pointer", "/data/attributes/name", "invalid")

		subject := errs[0]

		assert.Equal(t, subject.Title, "invalid")
		assert.Equal(t, subject.Detail, "Is not valid")
	})

	t.Run("when `title` is `not found`", func(t *testing.T) {
		errs := jsonapi.NewErrors("parameter", "id", "not found")

		subject := errs[0]

		assert.Equal(t, subject.Title, "not found")
		assert.Equal(t, subject.Detail, "Resource not found")
		assert.Equal(t, subject.Source.Parameter, "id")
	})

	t.Run("when `title` is `unknown`", func(t *testing.T) {
		errs := jsonapi.NewErrors("pointer", "/data/attributes/name", "unknown")

		subject := errs[0]

		assert.Equal(t, subject.Title, "unknown")
		assert.Equal(t, subject.Detail, "")
	})

	t.Run("when `detail` is specified", func(t *testing.T) {
		errs := jsonapi.NewErrors("pointer", "/data/attributes/name", "required", "Please provide your name")

		subject := errs[0]

		assert.Equal(t, subject.Title, "required")
		assert.Equal(t, subject.Detail, "Please provide your name")
	})
}

func TestMarshalJSON(t *testing.T) {
	errs := jsonapi.NewErrors("pointer", "/data/attributes/name", "required")

	expected := `{"errors":[{"title":"required","detail":"Can't be blank","source":{"pointer":"/data/attributes/name"}}]}`

	result, _ := errs.MarshalJSON()

	assert.Equal(t, string(result), expected)
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		errs := jsonapi.NewErrors()

		assert.Assert(t, errs.IsEmpty())
	})

	t.Run("when it is not empty", func(t *testing.T) {
		errs := jsonapi.NewErrors("pointer", "/data/attributes/name", "required")

		assert.Assert(t, !errs.IsEmpty())
	})
}
