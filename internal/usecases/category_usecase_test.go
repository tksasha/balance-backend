package usecases_test

import (
	"testing"

	"github.com/tksasha/balance/internal/interfaces/sqlite3/dummy"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/usecases"
	"gotest.tools/v3/assert"
)

func TestCreateCategory(t *testing.T) {
	params := &usecases.CategoryParams{Name: "Category #1"}

	t.Run("when Category should be created", func(t *testing.T) {
		expected := models.Category{ID: 2252, Name: "Category #1"}

		category, err := usecases.NewCategoryUsecase(
			dummy.CategoryRepository{
				ID: 2252,
			},
		).Create(params)

		assert.NilError(t, err)
		assert.Equal(t, *category, expected)
	})

	t.Run("when Category shouldn't be created", func(t *testing.T) {
		category, err := usecases.NewCategoryUsecase(
			dummy.CategoryRepository{},
		).Create(params)

		assert.Error(t, err, `{"application":{"database":"unknown"}}`)
		assert.Assert(t, category == nil)
	})
}

func TestFind(t *testing.T) {
	t.Run("when Category has found by ID", func(t *testing.T) {
		expected := models.BuildCategory("Category #2")

		category, err := usecases.NewCategoryUsecase(
			dummy.CategoryRepository{
				Category: expected,
			},
		).Find(1229)

		assert.NilError(t, err)
		assert.Equal(t, *category, *expected)
	})

	t.Run("when Category hasn't found by ID", func(t *testing.T) {
		category, err := usecases.NewCategoryUsecase(
			dummy.CategoryRepository{},
		).Find(1230)

		assert.Error(t, err, `{"application":{"database":"not found"}}`)
		assert.Assert(t, category == nil)
	})
}
