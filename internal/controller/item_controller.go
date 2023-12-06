package controller

import (
	"database/sql"
	"encoding/json"
	"os"

	sqlite3 "github.com/tksasha/balance/internal/interface/sqlite3/item"
	"github.com/tksasha/balance/internal/repository"
	"github.com/tksasha/balance/internal/usecase"
	"github.com/valyala/fasthttp"
)

type itemController struct {
	repo repository.ItemRepository
}

func NewItemController(db *sql.DB) *itemController {
	repo := sqlite3.New(db)

	return &itemController{repo}
}

func (controller *itemController) Show(ctx *fasthttp.RequestCtx) {
	uc := usecase.NewItemUsecase(controller.repo)

	item, _ := uc.Show(1)

	json.NewEncoder(os.Stdout).Encode(item)
}
