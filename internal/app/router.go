package app

import (
	"net/http"

	"github.com/tksasha/balance/internal/interfaces/api"
	"github.com/tksasha/balance/internal/interfaces/sqlite3"
)

func NewRouter(app *App) *http.ServeMux {
	mux := http.NewServeMux()

	repository := sqlite3.NewCategoryRepository(app.DBConn)

	categories := api.NewCategoryHandler(repository)

	mux.
		HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				categories.Create(w, r)

				return
			}
		})

	return mux
}
