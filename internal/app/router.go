package app

import (
	"github.com/gorilla/mux"
	"github.com/tksasha/balance/internal/interfaces/api"
	"github.com/tksasha/balance/internal/interfaces/sqlite3"
)

func NewRouter(app *App) *mux.Router {
	router := mux.NewRouter()

	repository := sqlite3.NewCategoryRepository(app.DBConn)

	router.
		HandleFunc("/categories", api.NewCategoryHandler(repository).Create).
		Methods("POST")

	return router
}
