package models

type Item struct {
	ID          int     `json:"id"`
	Formula     string  `json:"formula"`
	Sum         float64 `json:"sum"`
	CategoryID  int     `json:"category_id"`
	Description string  `json:"description"`
}

type ItemParams struct {
	Formula     string
	CategoryID  int
	Description string
}

func NewItem() *Item {
	return &Item{}
}
