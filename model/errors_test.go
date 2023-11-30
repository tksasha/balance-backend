package model_test

import (
	"testing"

	. "github.com/tksasha/balance/assert"
	"github.com/tksasha/balance/model"
)

func TestNewErrors(t *testing.T) {
	t.Run("when no one argument is provided", func(t *testing.T) {
		subject := model.NewErrors()

		Eq(t, subject.IsEmpty(), true)
	})

	t.Run("when only one argument is provided", func(t *testing.T) {
		subject := model.NewErrors("name")

		Eq(t, subject.IsEmpty(), true)
	})

	t.Run("when two arguments are provided", func(t *testing.T) {
		subject := model.NewErrors("name", "can't be blank")

		In(t, subject["errors"]["name"], "can't be blank")
	})
}

func TestAdd(t *testing.T) {
	t.Run("when attribute in lowercase", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("name", "required")

		In(t, subject["errors"]["name"], "required")
	})

	t.Run("when attribute in uppercase", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("Name", "required")

		In(t, subject["errors"]["name"], "required")
	})
}

func TestGet(t *testing.T) {
	subject := model.NewErrors()

	subject.Add("name", "can't be blank")

	In(t, subject["errors"]["name"], "can't be blank")
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		subject := model.NewErrors()

		Eq(t, subject.IsEmpty(), true)
	})

	t.Run("when it is not empty", func(t *testing.T) {
		subject := model.NewErrors()

		subject.Add("name", "can't be blank")

		Eq(t, subject.IsEmpty(), false)
	})
}
