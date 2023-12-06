package usecase_test

import (
	"gotest.tools/v3/assert"
	"testing"

	repository "github.com/tksasha/balance/internal/interface/testdb"
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/usecase"
)

func TestShow(t *testing.T) {
	repo := repository.New(nil)
	item := usecase.NewItemUsecase(repo)

	t.Run("when Item is found", func(t *testing.T) {
		exp := model.Item{
			ID:      42,
			Formula: "42.1 + 69.01",
			Sum:     111.11,
		}

		res, err := item.Show(42)

		assert.NilError(t, err)
		assert.Equal(t, *res, exp)
	})

	t.Run("when Item not found", func(t *testing.T) {
		res, err := item.Show(69)

		assert.Error(t, err, "Not Found")
		assert.Assert(t, res == nil)
	})
}
