package test

import (
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (repo *CategoryRepository) Create(*models.Category) error {
	return nil
}
