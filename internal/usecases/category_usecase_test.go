package usecases_test

import (
	"testing"

	repos "github.com/tksasha/balance/internal/interfaces/test"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/usecases"
	"gotest.tools/v3/assert"
)

var (
	categoryRepository = repos.NewCategoryRepository()
	categoryUsecase    = usecases.NewCategoryUsecase(categoryRepository)
)

func TestCreate(t *testing.T) {
	params := &usecases.CategoryParams{Name: "Category #1"}

	t.Run("when Category should be created", func(t *testing.T) {
		categoryRepository.
			SetCategoryShouldBeCreated(true).
			SetCategoryID(2252)

		category, err := categoryUsecase.Create(params)

		expected := models.Category{ID: 2252, Name: "Category #1"}

		assert.NilError(t, err)
		assert.Equal(t, *category, expected)
	})

	t.Run("when Category shouldn't be created", func(t *testing.T) {
		categoryRepository.
			SetCategoryShouldBeCreated(false)

		category, err := categoryUsecase.Create(params)

		assert.Error(t, err, "[APP ERROR] UNKNOWN")
		assert.Assert(t, category == nil)
	})
}

func TestShow(t *testing.T) {
	t.Run("when Category has found by ID", func(t *testing.T) {
		expected := models.BuildCategory("Category #2")

		categoryRepository.SetCategory(expected)

		category, err := categoryUsecase.Show(1809)

		assert.NilError(t, err)
		assert.Equal(t, *category, *expected)
	})

	t.Run("when Category hasn't found by ID", func(t *testing.T) {
		categoryRepository.SetCategory(nil)

		category, err := categoryUsecase.Show(1817)

		assert.Error(t, err, "[APP ERROR] NOT FOUND")
		assert.Assert(t, category == nil)
	})
}
