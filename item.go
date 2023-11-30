package main

import (
	"database/sql"
	"errors"
	"time"

	"github.com/tksasha/balance/date"
	"github.com/tksasha/balance/formula"
	"github.com/tksasha/balance/model"
)

type Item struct {
	model.Model
	ID          int       `json:"id"`
	Date        date.Date `json:"date"        validate:"required"`
	Formula     string    `json:"formula"     validate:"required"`
	Sum         float64   `json:"sum"`
	CategoryID  int       `json:"category_id" validate:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type itemParams struct {
	Date        date.Date `json:"date"`
	CategoryID  int       `json:"category_id"`
	Formula     string    `json:"formula"`
	Description string    `json:"description"`
}

func NewItem() *Item {
	item := &Item{}

	item.Errors = model.NewErrors()

	return item
}

func BuildItem(Date date.Date, Formula string, CategoryID int, Description string) *Item {
	item := &Item{
		Date:        Date,
		Formula:     Formula,
		CategoryID:  CategoryID,
		Description: Description,
	}

	item.Errors = model.NewErrors()

	return item
}

func (item *Item) calculate() {
	if len(item.Formula) == 0 {
		return
	}

	var err error

	item.Sum, err = formula.Calculate(item.Formula)
	if err != nil {
		item.Errors.Add("formula", "is not valid") // TODO: move it to constant
	}
}

func CreateItem(db *sql.DB, params *itemParams) (*Item, error) {
	item := BuildItem(
		params.Date,
		params.Formula,
		params.CategoryID,
		params.Description,
	)

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	item.calculate()

	model.Validate(item)
	if !item.IsValid() {
		return item, RecordInvalidError
	}

	sql := `
		INSERT INTO
			items(date, formula, sum, category_id, description, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?)
	`

	res, err := db.Exec(
		sql,
		item.Date.String(),
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.Description,
		item.CreatedAt,
		item.UpdatedAt,
	)
	if err != nil {
		return nil, InternalServerError // TODO: provide more details (Foreign Key Error)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, InternalServerError
	}

	item.ID = int(id)

	return item, nil
}

func FindItem(db *sql.DB, id int) (*Item, error) {
	item := new(Item)

	query := `
		SELECT
			id, date, formula, sum, category_id, description, created_at, updated_at
		FROM
			items
		WHERE
			id = ?
	`

	row := db.QueryRow(query, id)

	err := row.Scan(
		&item.ID,
		&item.Date,
		&item.Formula,
		&item.Sum,
		&item.CategoryID,
		&item.Description,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err == nil {
		return item, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, RecordNotFoundError
	}

	return nil, InternalServerError
}

func (item *Item) Update(db *sql.DB, params *itemParams) error {
	item.Date = params.Date
	item.Formula = params.Formula
	item.CategoryID = params.CategoryID
	item.Description = params.Description

	item.calculate()

	if model.Validate(item); !item.IsValid() {
		return RecordInvalidError
	}

	query := `
		UPDATE items
		SET date = ?, formula = ?, sum = ?, category_id = ?, description = ?
		WHERE id = ?
	`

	_, err := db.Exec(query, item.Date.String(), item.Formula, item.Sum, item.CategoryID, item.Description, item.ID)

	if err == nil {
		return nil
	}

	return InternalServerError
}
