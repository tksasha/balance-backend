package test

import (
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository struct {
	Category        *models.Category
	ShouldBeCreated bool
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (repo *CategoryRepository) SetCategory(category *models.Category) {
	repo.Category = category
}

func (repo *CategoryRepository) SetShouldBeCreated(flag bool) {
	repo.ShouldBeCreated = flag
}

func (repo *CategoryRepository) Create(*models.Category) error {
	if repo.ShouldBeCreated {
		return nil
	}

	return ErrUnknown
}

func (repo *CategoryRepository) Find(int) (*models.Category, error) {
	if repo.Category == nil {
		return nil, ErrNotFound
	}

	return repo.Category, nil
}
