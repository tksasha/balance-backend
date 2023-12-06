package category

import (
	"github.com/tksasha/balance/internal/interface/testdb/errors"
	"github.com/tksasha/balance/internal/model"
)

func (repo repository) Create(category *model.Category) error {
	if category.Name == "should be created" {
		category.ID = 1720

		return nil
	}

	return errors.UnknownDBError
}
