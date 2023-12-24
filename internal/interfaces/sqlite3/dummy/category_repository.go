package dummy

import (
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository struct {
	ID       int
	Category *models.Category
}

func (repository CategoryRepository) Create(category *models.Category) error {
	if repository.ID != 0 {
		category.ID = repository.ID

		return nil
	}

	return ErrUnknown
}

func (repository CategoryRepository) Find(int) (*models.Category, error) {
	if repository.Category == nil {
		return nil, ErrNotFound
	}

	return repository.Category, nil
}
