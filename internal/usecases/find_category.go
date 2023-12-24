package usecases

import (
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type FindCategory struct {
	repository repositories.CategoryFinder
}

func NewFindCategory(repository repositories.CategoryFinder) *FindCategory {
	return &FindCategory{repository}
}

func (usecase *FindCategory) Find(id int) (*models.Category, error) {
	category, err := usecase.repository.Find(id)
	if err != nil {
		return nil, NewNotFoundError(err)
	}

	return category, nil
}
