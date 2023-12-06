package usecase

import (
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/repository"
	"github.com/tksasha/balance/internal/usecase/errors"
	validator "github.com/tksasha/balance/pkg/model"
)

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{repo}
}

func (uc *categoryUsecase) Create(params *model.CategoryParams) (*model.Category, error) {
	category := model.BuildCategory(params.Name)

	if validator.Validate(category); !category.IsValid() {
		return category, errors.RecordInvalidError
	}

	if err := uc.repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}
