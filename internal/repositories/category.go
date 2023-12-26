package repositories

import (
	"github.com/tksasha/balance/internal/models"
)

type Category interface {
	Create(category *models.Category) error
	Find(id int) (*models.Category, error)
}
