package app

import (
	"log"
	"net/http"

	"github.com/tksasha/balance/config"
	"github.com/tksasha/balance/internal/app/router"
)

func Run(config *config.Config) {
	handler := router.New(config)

	log.Fatal(http.ListenAndServe(":3000", handler))
}
