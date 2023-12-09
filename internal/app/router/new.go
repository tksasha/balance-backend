package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tksasha/balance/config"
	controllers "github.com/tksasha/balance/internal/interfaces/api"
	database "github.com/tksasha/balance/internal/interfaces/sqlite3"
)

func New(config *config.Config) *mux.Router {
	router := mux.NewRouter()
	repository := database.Open(config)

	// items := controller.NewItemController(db)
	// r.HandleFunc("/items/{id}", items.Show).Methods(http.MethodGet)

	categories := controllers.NewCategoriesController(repository.NewCategoryRepository())
	router.
		HandleFunc("/categories", categories.Create).
		Methods(http.MethodPost)

	return router
}
