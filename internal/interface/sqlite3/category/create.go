package category

import (
	"context"

	"github.com/tksasha/balance/internal/model"
)

func (repo repository) Create(category *model.Category) error {
	query := `INSERT INTO categories(name) VALUES(?)`

	res, err := repo.db.ExecContext(context.TODO(), query, category.Name)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = int(id)

	return nil
}
