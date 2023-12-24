package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/usecases"
)

type CategoryHandler struct {
	repository repositories.CategoryCreator
}

func NewCategoryHandler(repository repositories.CategoryCreator) *CategoryHandler {
	return &CategoryHandler{repository}
}

func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err) // TODO: return Bad Request
	}

	params := new(usecases.CategoryParams)

	if err := json.Unmarshal(data, params); err != nil {
		panic(err) // TODO: return Bad Request
	}

	category, err := usecases.
		NewCreateCategory(handler.repository).
		Create(params)

	if err != nil {
		JSON(w, err.Error())

		return
	}

	w.Header().Set("content-type", "application/json")

	JSON(w, category)
}
