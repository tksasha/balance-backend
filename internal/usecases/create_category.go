package usecases

import (
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type CreateCategory struct {
	repository repositories.CategoryCreator
}

type CategoryParams struct {
	Name string `json:"name"`
}

func NewCreateCategory(repository repositories.CategoryCreator) *CreateCategory {
	return &CreateCategory{repository}
}

func (usecase *CreateCategory) Create(params *CategoryParams) (*models.Category, error) {
	category := models.BuildCategory(params.Name)

	if err := usecase.repository.Create(category); err != nil {
		return nil, NewUnknownError(err)
	}

	return category, nil
}
