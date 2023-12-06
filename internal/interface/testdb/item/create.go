package item

import (
	"github.com/tksasha/balance/internal/interface/testdb/errors"
	"github.com/tksasha/balance/internal/model"
)

func (repo repository) Create(item *model.Item) error {
	if item.CategoryID == 42 {
		item.ID = 1610

		return nil
	}

	return errors.UnknownDBError
}
