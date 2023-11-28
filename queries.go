package main

import (
	"database/sql"
	"errors"
)

func CreateItemQuery(db *sql.DB, item *Item) error {
	sql := `
		INSERT INTO
			items(date, formula, sum, category_id, description)
		VALUES
			(?, ?, ?, ?, ?)
	`

	res, err := db.Exec(sql, item.Date.String(), item.Formula, item.Sum, item.CategoryID, item.Description) // TODO: do not forget about timestamps
	if err != nil {
		return err // TODO: prettify this error
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err // TODO: prettify this error
	}

	item.ID = id

	return nil
}

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
