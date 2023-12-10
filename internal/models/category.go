package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewCategory() *Category {
	return &Category{}
}

func BuildCategory(name string) *Category {
	return &Category{Name: name}
}
