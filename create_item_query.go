package main

import (
	"database/sql"
	"errors"

	"github.com/tksasha/model"
)

func CreateItemQuery(db *sql.DB, item *Item) error {
	sql := `
		INSERT INTO
			items(date, formula, sum, category_id, description)
		VALUES
			(?, ?, ?, ?, ?)
	`

	res, err := db.Exec(sql, item.Date.String(), item.Formula, item.Sum, item.CategoryID, item.Description)
	if err != nil {
		return errors.New(model.ErrExecSQL)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.New(model.ErrObtainID)
	}

	item.ID = id

	return nil
}
