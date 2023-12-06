package sqlite3

import (
	"context"

	"github.com/tksasha/balance/internal/model"
)

func (repo *SQLite3Repository) FindByID(id int) (*model.Item, error) {
	item := model.NewItem()

	query := `SELECT id, formula, sum FROM items WHERE id = ?`

	row := repo.DB.QueryRowContext(context.TODO(), query, id)

	if err := row.Scan(&item.ID, &item.Formula, &item.Sum); err != nil {
		// TODO: sql.ErrNoRows -> ErrNotFound otherwise ErrUnknown
		return nil, err
	}

	return item, nil
}