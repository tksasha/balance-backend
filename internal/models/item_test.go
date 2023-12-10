package models_test

import (
	"testing"

	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestNewItem(t *testing.T) {
	exp := models.Item{Formula: "2+2"}

	sbj := models.NewItem()

	assert.Equal(t, exp, *sbj)
}
