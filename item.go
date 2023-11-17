package main

import (
	"github.com/tksasha/model"
	"github.com/tksasha/validations"
	"github.com/tksasha/date"
)

type Item struct {
	model.Model
	ID          int       `json:"id"`
	Date        date.Date `json:"date"`
	Sum         float32   `json:"sum"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	Formula     string    `json:"formula"`
}

func NewItem() (item *Item) {
	item = &Item{}

	item.Errors = validations.NewErrors()

	return
}

func (item *Item) Validate() {
	if item.Date.IsEmpty() {
		item.Errors.Add("date", validations.BLANK)
	}

	if item.CategoryID == 0 {
		item.Errors.Add("category_id", validations.BLANK)
	}

	validations.ValidatePresenceOf(item.Errors, "formula", item.Formula)
}
