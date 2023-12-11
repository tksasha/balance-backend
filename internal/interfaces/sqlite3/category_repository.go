package sqlite3

import (
	"database/sql"
	"fmt"

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
		return fmt.Errorf("[DB][QUERY] %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("[DB][ID] %w", err)
	}

	category.ID = int(id)

	return nil
}

func (repository *CategoryRepository) Find(int) (*models.Category, error) {
	return &models.Category{}, nil
}
