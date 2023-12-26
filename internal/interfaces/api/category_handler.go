package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/usecases"
)

type CategoryHandler struct {
	usecase *usecases.CategoryUsecase
}

func NewCategoryHandler(repository repositories.Category) *CategoryHandler {
	return &CategoryHandler{
		usecases.NewCategoryUsecase(repository),
	}
}

func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		JSON(w, NewError(err))

		return
	}

	params := new(usecases.CategoryParams)

	if err := json.Unmarshal(data, params); err != nil {
		JSON(w, NewError(err))

		return
	}

	category, err := handler.usecase.Create(params)

	if err != nil {
		JSON(w, NewError(err))

		return
	}

	JSON(w, category)
}
