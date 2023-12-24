package sqlite3

import (
	"database/sql"

	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (repository *CategoryRepository) Create(category *models.Category) error {
	query := `INSERT INTO categories(name) VALUES(?)`

	res, err := repository.db.Exec(query, category.Name)
	if err != nil {
		return NewQueryError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return NewGetIDError(err)
	}

	category.ID = int(id)

	return nil
}

func (repository *CategoryRepository) Find(int) (*models.Category, error) {
	return &models.Category{}, nil
}
