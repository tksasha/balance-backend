package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tksasha/balance/internal/controller"
)

func New(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	items := controller.NewItemController(db)
	r.HandleFunc("/items/{id}", items.Show).Methods(http.MethodGet)

	categories := controller.NewCategoryController(db)
	r.HandleFunc("/categories", categories.Create).Methods(http.MethodPost)

	return r
}
