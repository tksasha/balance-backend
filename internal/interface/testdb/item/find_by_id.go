package item

import (
	"errors"

	"github.com/tksasha/balance/internal/model"
)

func (repo repository) FindByID(id int) (*model.Item, error) {
	if id == 42 {
		return item, nil
	}

	return nil, errors.New("Not Found")
}
