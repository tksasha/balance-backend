package main

import (
	"time"

	"github.com/tksasha/model"
	"github.com/tksasha/validations"
)

type Item struct {
	model.Model
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Sum         float32   `json:"sum"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	Formula     string    `json:"formula"`
}

func NewItem() *Item {
	item := Item{}

	item.Errors = validations.NewErrors()

	return &item
}

func (item *Item) Validate() {
	if item.Date == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		item.Errors.Add("date", validations.CANT_BE_BLANK)
	}
}
