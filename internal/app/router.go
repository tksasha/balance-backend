package app

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controllers "github.com/tksasha/balance/internal/interfaces/api"
	"github.com/tksasha/balance/internal/interfaces/sqlite3"
)

func NewRouter(app *App) *mux.Router {
	router := mux.NewRouter()

	repository := sqlite3.NewCategoryRepository(app.DBConn)

	categories := controllers.NewCategoryController(repository)

	router.
		Handle("/categories", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categories.Create))).
		Methods("POST")

	return router
}
