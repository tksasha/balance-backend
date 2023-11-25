package main

import (
	"database/sql"

	"github.com/tksasha/balance/date"
	"github.com/tksasha/balance/formula"
	"github.com/tksasha/balance/model"
)

type Item struct {
	model.Model
	ID          int64     `json:"id"`
	Date        date.Date `json:"date"        validate:"required"`
	Formula     string    `json:"formula"     validate:"required"`
	Sum         float64   `json:"sum"`
	CategoryID  int       `json:"category_id" validate:"required"`
	Description string    `json:"description"`
}

func NewItem() (item *Item) {
	item = new(Item)

	item.Errors = model.NewErrors()

	return
}

func (item *Item) Calculate() {
	if len(item.Formula) > 0 {
		sum := 0.0

		var err error

		if sum, err = formula.Calculate(item.Formula); err != nil {
			item.Errors.Add("formula", "is not valid")
		}

		item.Sum = sum
	}
}

func CreateItem(db *sql.DB, params *ItemParams) (*Item, error) {
	item := NewItem()

	item.Date = params.Date
	item.CategoryID = params.CategoryID
	item.Formula = params.Formula
	item.Description = params.Description

	item.Calculate()

	model.Validate(item)

	if item.IsNotValid() {
		return item, ClientError
	}

	if err := CreateItemQuery(db, item); err != nil {
		return nil, err
	}

	return item, nil
}

// func FindItem(db *sql.DB, id int64) *Item {
// 	item := NewItem(db)

// 	row := db.QueryRow(`SELECT id, date FROM items WHERE id = ?`, id)

// 	if err := row.Scan(&item.ID, &item.Date); err != nil {
// 		// panic(err)
// 		return nil
// 	}

// 	return item
// }
