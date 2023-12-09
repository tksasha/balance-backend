package item

import (
	"github.com/tksasha/balance/internal/interfaces/testdb/errors"
	"github.com/tksasha/balance/internal/models"
)

func (repo repository) Create(item *models.Item) error {
	if item.CategoryID == 42 {
		item.ID = 1610

		return nil
	}

	return errors.UnknownDBError
}
