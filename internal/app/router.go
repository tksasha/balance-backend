package app

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controllers "github.com/tksasha/balance/internal/interfaces/api"
	"github.com/tksasha/balance/internal/interfaces/dummydb"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	repository := dummydb.NewCategoryRepository()

	categories := controllers.NewCategoryController(repository)

	router.
		Handle("/categories", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categories.Create))).
		Methods("POST")

	return router
}
