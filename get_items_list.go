package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetItemsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sql := `
		SELECT
			id,
			date,
			sum,
			category_id,
			description,
			formula
		FROM
			items
		ORDER BY
			created_at DESC
		LIMIT
			10
	`

	rows, err := DB.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var Items []Item

	for rows.Next() {
		item := Item{}

		err = rows.Scan(&item.ID, &item.Date, &item.Sum, &item.CategoryID, &item.Description, &item.Formula)
		if err != nil {
			log.Fatal(err)
		}

		Items = append(Items, item)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	OK(w, Items)
}
