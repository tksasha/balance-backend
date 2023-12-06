package app

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/config"
	"github.com/tksasha/balance/internal/app/db"
	"github.com/tksasha/balance/internal/app/router"
	"github.com/valyala/fasthttp"
)

func Run(c *config.Config) {
	conn := db.Open(c)

	if conn != nil {
		defer db.Close(conn)
	}

	log.Fatal(fasthttp.ListenAndServe(":3000", router.New()))
}
