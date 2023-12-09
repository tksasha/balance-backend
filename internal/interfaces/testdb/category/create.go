package category

import (
	"github.com/tksasha/balance/internal/interfaces/testdb/errors"
	"github.com/tksasha/balance/internal/models"
)

func (repo repository) Create(category *models.Category) error {
	if category.Name == "should be created" {
		category.ID = 1720

		return nil
	}

	return errors.UnknownDBError
}
