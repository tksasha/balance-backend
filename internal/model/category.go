package model

import (
	"github.com/tksasha/balance/pkg/model"
	"github.com/tksasha/balance/pkg/model/errors"
)

type Category struct {
	model.Model
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

type CategoryParams struct {
	Name string `json:"name"`
}

func NewCategory() *Category {
	category := &Category{}

	category.Errors = errors.New()

	return category
}

func BuildCategory(name string) *Category {
	category := NewCategory()

	category.Name = name

	return category
}
