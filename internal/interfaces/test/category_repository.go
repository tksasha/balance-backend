package test

import (
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository struct {
	Category                *models.Category
	CategoryID              int
	CategoryShouldBeCreated bool
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (repository *CategoryRepository) SetCategory(category *models.Category) *CategoryRepository {
	repository.Category = category

	return repository
}

func (repository *CategoryRepository) SetCategoryID(id int) *CategoryRepository {
	repository.CategoryID = id

	return repository
}

func (repository *CategoryRepository) SetCategoryShouldBeCreated(flag bool) *CategoryRepository {
	repository.CategoryShouldBeCreated = flag

	return repository
}

func (repository *CategoryRepository) Create(category *models.Category) error {
	if repository.CategoryShouldBeCreated {
		category.ID = repository.CategoryID

		return nil
	}

	return ErrUnknown
}

func (repository *CategoryRepository) Find(int) (*models.Category, error) {
	if repository.Category == nil {
		return nil, ErrNotFound
	}

	return repository.Category, nil
}
