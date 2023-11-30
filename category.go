package main

import (
	"database/sql"
	"errors"

	"github.com/tksasha/balance/model"
)

type Category struct {
	model.Model
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

type categoryParams struct {
	Name string
}

func NewCategory() *Category {
	return &Category{model.Model{Errors: model.NewErrors()}, 0, ""}
}

func BuildCategory(name string) *Category {
	return &Category{model.Model{Errors: model.NewErrors()}, 0, name}
}

func CreateCategory(db *sql.DB, params *categoryParams) (*Category, error) {
	category := BuildCategory(params.Name)

	model.Validate(category)
	if !category.IsValid() {
		return category, RecordInvalidError
	}

	query := `INSERT INTO categories(name) VALUES (?)`

	// TODO: do not forget about timestamps
	res, err := db.Exec(query, category.Name)
	if err != nil {
		return nil, InternalServerError
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, InternalServerError
	}

	category.ID = int(id)

	return category, nil
}

func FindCategory(db *sql.DB, id int) (*Category, error) {
	category := NewCategory()

	query := `SELECT id, name FROM categories WHERE id = ?`

	row := db.QueryRow(query, id)

	err := row.Scan(&category.ID, &category.Name)

	if err == nil {
		return category, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, RecordNotFoundError
	}

	return nil, InternalServerError
}
