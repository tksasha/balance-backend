package repository

import (
	"github.com/tksasha/balance/internal/model"
)

type ItemRepository interface {
	FindByID(int) (*model.Item, error)
}
