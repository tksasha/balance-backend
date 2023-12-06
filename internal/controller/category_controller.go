package controller

import (
	"database/sql"
	"encoding/json"

	sqlite3 "github.com/tksasha/balance/internal/interface/sqlite3/category"
	"github.com/tksasha/balance/internal/model"
	"github.com/tksasha/balance/internal/repository"
	"github.com/tksasha/balance/internal/usecase"
	"github.com/valyala/fasthttp"
)

type categoryController struct {
	repo repository.CategoryRepository
}

func NewCategoryController(db *sql.DB) *categoryController {
	repo := sqlite3.New(db)

	return &categoryController{repo}
}

func (controller *categoryController) Create(ctx *fasthttp.RequestCtx) {
	categories := usecase.NewCategoryUsecase(controller.repo)

	params := &model.CategoryParams{}

	json.Unmarshal(ctx.PostBody(), params)

	category, err := categories.Create(params)
	if err != nil {
		json.NewEncoder(ctx).Encode(err)

		return
	}

	ctx.SetStatusCode(201) // Created

	json.NewEncoder(ctx).Encode(category)
}
