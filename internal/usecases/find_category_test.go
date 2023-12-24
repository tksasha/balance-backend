package usecases_test

import (
	"testing"

	"github.com/tksasha/balance/internal/interfaces/sqlite3/dummy"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/usecases"
	"gotest.tools/v3/assert"
)

func TestFind(t *testing.T) {
	t.Run("when Category has found by ID", func(t *testing.T) {
		expected := models.BuildCategory("Category #2")

		category, err := usecases.NewFindCategory(
			dummy.CategoryRepository{
				Category: expected,
			},
		).Find(1229)

		assert.NilError(t, err)
		assert.Equal(t, *category, *expected)
	})

	t.Run("when Category hasn't found by ID", func(t *testing.T) {
		category, err := usecases.NewFindCategory(
			dummy.CategoryRepository{},
		).Find(1230)

		assert.Error(t, err, `{"application":{"NotFoundError":"NOT FOUND"}}`)
		assert.Assert(t, category == nil)
	})
}
