package models

type ItemRepository interface {
	Find(id int) (*Item, error)
	Create(*Item) error
}
