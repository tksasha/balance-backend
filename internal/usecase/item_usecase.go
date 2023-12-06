package usecase

import (
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/repository"
	"github.com/tksasha/balance/internal/usecase/errors"
)

type itemUsecase struct {
	repo repository.ItemRepository
}

func NewItemUsecase(repo repository.ItemRepository) *itemUsecase {
	return &itemUsecase{repo}
}

func (uc *itemUsecase) Show(id int) (*model.Item, error) {
	return uc.repo.FindByID(id)
}

func (uc *itemUsecase) Create(params model.ItemParams) (*model.Item, error) {
	item := model.NewItem()

	item.Formula = params.Formula
	item.CategoryID = params.CategoryID
	item.Description = params.Description

	err := uc.repo.Create(item)
	if err != nil {
		return item, errors.RecordInvalidError
	}

	return item, nil
}
