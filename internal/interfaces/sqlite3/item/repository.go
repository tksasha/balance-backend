package item

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

func (repo repository) Create(item *models.Item) error {
	return nil
}

func (repo repository) Find(id int) (*models.Item, error) {
	return nil, nil
}
