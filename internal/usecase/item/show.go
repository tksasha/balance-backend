package item

import (
	"github.com/tksasha/balance/internal/model"
)

func (usecase *itemUsecase) Show(id int) (*model.Item, error) {
	return usecase.repo.FindByID(id)
}
