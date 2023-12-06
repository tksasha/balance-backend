package item

import (
	"github.com/tksasha/balance/internal/repository"
)

type itemUsecase struct {
	repo repository.ItemRepository
}

func New(repo repository.ItemRepository) *itemUsecase {
	return &itemUsecase{repo}
}
