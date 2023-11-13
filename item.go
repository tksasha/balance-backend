package main

import (
	"time"
)

type Item struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Sum         float32   `json:"sum"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	Formula     string    `json:"formula"`
}
