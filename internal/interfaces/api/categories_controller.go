package api

import (
	"net/http"

	"github.com/tksasha/balance/internal/models"
	usecase "github.com/tksasha/balance/internal/usecases/category"
)

type categoriesController struct {
	usecase *usecase.Usecase
}

func NewCategoriesController(repository models.CategoryRepository) *categoriesController {
	return &categoriesController{usecase.New(repository)}
}

func (controller categoriesController) Create(w http.ResponseWriter, r *http.Request) {
}
