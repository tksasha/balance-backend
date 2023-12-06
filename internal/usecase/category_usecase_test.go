package usecase_test

import (
	"testing"

	testdb "github.com/tksasha/balance/internal/interface/testdb/category"
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/usecase"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCreateCategory(t *testing.T) {
	repo := testdb.New()
	subject := usecase.NewCategoryUsecase(repo)

	t.Run("when params are not valid", func(t *testing.T) {
		params := &model.CategoryParams{}

		category, err := subject.Create(params)

		errs := category.Errors["errors"]

		assert.Error(t, err, "Record Invalid Error")
		assert.Assert(t, is.Contains(errs["name"], "required"))
	})

	t.Run("when something happened with database", func(t *testing.T) {
		params := &model.CategoryParams{Name: "to raise db error"}

		category, err := subject.Create(params)

		assert.Error(t, err, "Unknown DB Error")
		assert.Assert(t, is.Nil(category))
	})

	t.Run("when params are valid", func(t *testing.T) {
		params := &model.CategoryParams{Name: "should be created"}

		category, err := subject.Create(params)

		assert.NilError(t, err)
		assert.Equal(t, category.ID, 1720)
		assert.Equal(t, category.Name, "should be created")
	})
}
