package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/tksasha/balance/config"
)

type App struct {
	DBConn *sql.DB
}

func Run() {
	config := config.NewConfig()

	app := &App{DBConn: OpenDB(config)}

	router := NewRouter(app)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}

	log.Fatal(server.ListenAndServe())
}
