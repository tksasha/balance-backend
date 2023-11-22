package main

import (
	"database/sql"
	"errors"

	"github.com/tksasha/date"
	"github.com/tksasha/formula"
	"github.com/tksasha/model"
)

type Item struct {
	model.Model
	ID          int64     `json:"id"`
	Date        date.Date `json:"date"`
	Formula     string    `json:"formula"`
	Sum         float64   `json:"sum"`
	CategoryID  int       `json:"category_id"`
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
			item.Errors.Add("formula", model.ErrIsNotValid)
		}

		item.Sum = sum
	}
}

func (item *Item) Validate() {
	if item.Date.IsEmpty() {
		item.Errors.Add("date", model.ErrIsBlank)
	}

	if item.CategoryID == 0 {
		item.Errors.Add("category_id", model.ErrIsBlank)
	}

	if item.Formula == "" {
		item.Errors.Add("formula", model.ErrIsBlank)
	}
}

func CreateItem(db *sql.DB, params *ItemParams) (*Item, error) {
	item := new(Item)

	item.Errors = model.NewErrors()

	item.Date = params.Date
	item.CategoryID = params.CategoryID
	item.Formula = params.Formula
	item.Description = params.Description

	item.Calculate()

	item.Validate()

	if item.IsNotValid() {
		return item, errors.New(model.ErrIsNotValid)
	}

	sql := `
		INSERT INTO
			items(date, formula, sum, category_id, description)
		VALUES
			(?, ?, ?, ?, ?)
	`

	res, err := db.Exec(sql, item.Date.String(), item.Formula, item.Sum, item.CategoryID, item.Description)
	if err != nil {
		return item, errors.New(model.ErrExecSQL)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return item, errors.New(model.ErrObtainID)
	}

	item.ID = id

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
