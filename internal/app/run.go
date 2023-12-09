package app

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/config"
	"github.com/tksasha/balance/internal/app/db"
	"github.com/tksasha/balance/internal/app/router"
)

func Run(c *config.Config) {
	conn := db.Open(c)
	if conn != nil {
		defer db.Close(conn)
	}

	handler := router.New(conn)

	log.Fatal(http.ListenAndServe(":3000", handler))
}
