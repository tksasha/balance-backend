package controller

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	sqlite3 "github.com/tksasha/balance/internal/interface/sqlite3/category"
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/repository"
	"github.com/tksasha/balance/internal/usecase"
)

type categoryController struct {
	repo repository.CategoryRepository
}

func NewCategoryController(db *sql.DB) *categoryController {
	repo := sqlite3.New(db)

	return &categoryController{repo}
}

func (controller *categoryController) Create(w http.ResponseWriter, r *http.Request) {
	categories := usecase.NewCategoryUsecase(controller.repo)

	params := &model.CategoryParams{}

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, params)

	category, err := categories.Create(params)
	if err != nil {
		json.NewEncoder(w).Encode(err)

		return
	}

	json.NewEncoder(w).Encode(category)
}
