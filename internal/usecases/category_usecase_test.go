package usecases_test

import (
	"testing"

	repos "github.com/tksasha/balance/internal/interfaces/test"
	"github.com/tksasha/balance/internal/usecases"
	"gotest.tools/v3/assert"
)

func TestCreate(t *testing.T) {
	repo := repos.NewCategoryRepository()

	usecase := usecases.NewCategoryUsecase(repo)

	params := &usecases.CategoryParams{Name: "Category #1"}

	err := usecase.Create(params)

	assert.NilError(t, err)
}
