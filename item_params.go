package main

import (
	"github.com/tksasha/balance/date"
)

type ItemParams struct {
	Date        date.Date `json:"date"`
	CategoryID  int       `json:"category_id"`
	Formula     string    `json:"formula"`
	Description string    `json:"description"`
}
