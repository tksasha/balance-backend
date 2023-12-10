package models

type Item struct {
	Formula string
}

func NewItem() *Item {
	return &Item{"2+2"}
}
