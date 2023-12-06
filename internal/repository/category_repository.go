package repository

import (
	"github.com/tksasha/balance/internal/model"
)

type CategoryRepository interface {
	Create(*model.Category) error
}
