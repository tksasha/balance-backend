package model

type Item struct {
	ID      int     `json:"id"`
	Formula string  `json:"formula"`
	Sum     float64 `json:"sum"`
}

func NewItem() *Item {
	return &Item{}
}
