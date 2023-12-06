package usecase

import (
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/repository"
)

type itemUsecase struct {
	repo repository.ItemRepository
}

func NewItemUsecase(repo repository.ItemRepository) *itemUsecase {
	return &itemUsecase{repo}
}

func (usecase *itemUsecase) Show(id int) (*model.Item, error) {
	return usecase.repo.FindByID(id)
}
