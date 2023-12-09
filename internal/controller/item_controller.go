package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	sqlite3 "github.com/tksasha/balance/internal/interface/sqlite3/item"
	"github.com/tksasha/balance/internal/repository"
	"github.com/tksasha/balance/internal/usecase"
)

type itemController struct {
	repo repository.ItemRepository
}

func NewItemController(db *sql.DB) *itemController {
	repo := sqlite3.New(db)

	return &itemController{repo}
}

func (controller *itemController) Show(w http.ResponseWriter, r *http.Request) {
	uc := usecase.NewItemUsecase(controller.repo)

	item, _ := uc.Show(1)

	json.NewEncoder(w).Encode(item)
}
