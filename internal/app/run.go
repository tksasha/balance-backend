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

	log.Fatal(http.ListenAndServe(":3000", router))
}
