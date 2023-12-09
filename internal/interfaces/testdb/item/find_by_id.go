package item

import (
	"errors"

	"github.com/tksasha/balance/internal/models"
)

func (repo repository) FindByID(id int) (*models.Item, error) {
	if id == 42 {
		return item, nil
	}

	return nil, errors.New("Not Found")
}
