package repositories

import (
	"github.com/tksasha/balance/internal/models"
)

type CategoryFinder interface {
	Find(id int) (*models.Category, error)
}
