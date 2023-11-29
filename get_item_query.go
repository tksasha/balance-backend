package main

import (
	"database/sql"
	"errors"
)

func GetItemQuery(db *sql.DB, id int) (*Item, error) {
	item := new(Item)

	query := `
		SELECT
			id, date, formula, sum, category_id, description
		FROM
			items
		WHERE
			id = ?
	`

	row := db.QueryRow(query, id)
	if err := row.Scan(&item.ID, &item.Date, &item.Formula, &item.Sum, &item.CategoryID, &item.Description); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, RecordNotFoundError
		}
	}

	return item, nil
}
