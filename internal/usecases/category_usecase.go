package usecases

import (
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type CategoryUsecase struct {
	repository repositories.CategoryRepository
}

type CategoryParams struct {
	Name string `json:"name"`
}

func NewCategoryUsecase(repository repositories.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repository}
}

func (usecase *CategoryUsecase) Create(params *CategoryParams) (*models.Category, error) {
	category := models.BuildCategory(params.Name)

	if err := usecase.repository.Create(category); err != nil {
		return nil, ErrUnknown
	}

	return category, nil
}

func (usecase *CategoryUsecase) Show(id int) (*models.Category, error) {
	category, err := usecase.repository.Find(id)
	if err != nil {
		return nil, ErrNotFound
	}

	return category, nil
}
