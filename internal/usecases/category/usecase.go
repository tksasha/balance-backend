package category

import (
	"github.com/tksasha/balance/internal/models"
)

type Usecase struct {
	repository models.CategoryRepository
}

func New(repository models.CategoryRepository) *Usecase {
	return &Usecase{repository}
}
