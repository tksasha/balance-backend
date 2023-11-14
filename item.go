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

func NewItem() (item *Item) {
	item = &Item{}

	item.Errors = validations.NewErrors()

	return
}

func (item *Item) Validate() {
	if item.Date == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		item.Errors.Add("date", validations.BLANK)
	}

	if item.CategoryID == 0 {
		item.Errors.Add("category_id", validations.BLANK)
	}

	validations.ValidatePresenceOf(item.Errors, "formula", item.Formula)
}
