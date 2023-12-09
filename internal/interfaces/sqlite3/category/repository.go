package category

import (
	"database/sql"

	"github.com/tksasha/balance/internal/models"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return repository{db}
}

func (repo repository) Create(category *models.Category) error {
	return nil
}

func (repo repository) Find(id int) (*models.Category, error) {
	return nil, nil
}
