package models

type CategoryRepository interface {
	Create(*Category) error
}
