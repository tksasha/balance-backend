package repositories

import (
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
}
