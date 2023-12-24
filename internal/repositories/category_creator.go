package repositories

import (
	"github.com/tksasha/balance/internal/models"
)

type CategoryCreator interface {
	Create(category *models.Category) error
}
