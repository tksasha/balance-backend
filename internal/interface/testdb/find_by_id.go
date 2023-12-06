package testdb

import (
	"errors"
	
	"github.com/tksasha/balance/internal/model"
)

var (
	item = &model.Item{
		ID: 42,
		Formula: "42.1 + 69.01",
		Sum: 111.11,
	}
)

func (repo repository) FindByID(id int) (*model.Item, error) {
	if id == 42 {
		return item, nil
	}

	return nil, errors.New("Not Found")
}
