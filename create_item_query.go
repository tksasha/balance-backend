package main

import (
	"database/sql"
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
