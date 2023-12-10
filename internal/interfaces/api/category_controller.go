package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/usecases"
)

type CategoryController struct {
	usecase *usecases.CategoryUsecase
}

func NewCategoryController(repository repositories.CategoryRepository) *CategoryController {
	usecase := usecases.NewCategoryUsecase(repository)

	return &CategoryController{usecase}
}

func (controller *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err) // TODO: return Bad Request
	}

	params := new(usecases.CategoryParams)

	if err := json.Unmarshal(data, params); err != nil {
		panic(err) // TODO: return Bad Request
	}

	category, err := controller.usecase.Create(params)
	if err != nil {
		panic(err)
	}

	JSON(w, category)
}
