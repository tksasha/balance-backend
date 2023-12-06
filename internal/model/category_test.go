package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/model"
	"gotest.tools/v3/assert"
)

func TestNewCategory(t *testing.T) {
	category := model.NewCategory()

	assert.Assert(t, category.IsValid())
}

func TestBuildCategory(t *testing.T) {
	category := model.BuildCategory("Category One")

	assert.Equal(t, category.Name, "Category One")
}
