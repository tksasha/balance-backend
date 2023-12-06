package usecase_test

import (
	"gotest.tools/v3/assert"
	"testing"

	testdb "github.com/tksasha/balance/internal/interface/testdb/item"
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/usecase"
)

func TestShowItem(t *testing.T) {
	repo := testdb.New()
	subj := usecase.NewItemUsecase(repo)

	t.Run("when Item is found", func(t *testing.T) {
		exp := model.Item{
			ID:      42,
			Formula: "42.1 + 69.01",
			Sum:     111.11,
		}

		res, err := subj.Show(42)

		assert.NilError(t, err)
		assert.Equal(t, *res, exp)
	})

	t.Run("when Item not found", func(t *testing.T) {
		res, err := subj.Show(69)

		assert.Error(t, err, "Not Found")
		assert.Assert(t, res == nil)
	})
}

func TestCreateItem(t *testing.T) {
	repo := testdb.New()
	subj := usecase.NewItemUsecase(repo)

	t.Run("when params are not valid", func(t *testing.T) {
		params := model.ItemParams{}

		res, err := subj.Create(params)

		assert.Error(t, err, "Record Invalid Error")
		assert.Assert(t, res.ID == 0)
	})

	t.Run("when params are valid", func(t *testing.T) {
		params := model.ItemParams{
			Formula:     "42.1 + 69.01",
			CategoryID:  42,
			Description: "lorem ipsum ...",
		}

		res, err := subj.Create(params)

		assert.NilError(t, err)
		assert.Assert(t, res.ID > 0)
	})
}
