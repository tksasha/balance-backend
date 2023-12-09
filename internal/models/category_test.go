package models_test

import (
	"testing"

	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestNewCategory(t *testing.T) {
	category := models.NewCategory()

	assert.Assert(t, category.IsValid())
}

func TestBuildCategory(t *testing.T) {
	category := models.BuildCategory("Category One")

	assert.Equal(t, category.Name, "Category One")
}
