package main

import (
	"database/sql"
)

func GetItemQuery(db *sql.DB, id int) (*Item, error) {
	item := new(Item)

	sql := `
		SELECT
			id, date, formula, sum, category_id, description
		FROM
			items
		WHERE
			id = ?
	`

	row := db.QueryRow(sql, id)
	if err := row.Scan(&item.ID, &item.Date, &item.Formula, &item.Sum, &item.CategoryID, &item.Description); err != nil {
		return item, err
	}

	return item, nil
}
