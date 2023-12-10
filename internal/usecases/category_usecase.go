package usecases

import (
	"github.com/tksasha/balance/internal/models"
	repos "github.com/tksasha/balance/internal/repositories"
)

type CategoryUsecase struct {
	repo repos.CategoryRepository
}

type CategoryParams struct {
	Name string
}

func NewCategoryUsecase(repo repos.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo}
}

func (usecase *CategoryUsecase) Create(params *CategoryParams) error {
	category := models.BuildCategory(params.Name)

	if err := usecase.repo.Create(category); err != nil {
		return ErrUnknown
	}

	return nil
}

func (usecase *CategoryUsecase) Show(id int) (*models.Category, error) {
	category, err := usecase.repo.Find(id)
	if err != nil {
		return nil, ErrNotFound
	}

	return category, nil
}
