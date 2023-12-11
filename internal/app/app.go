package app

import (
	"log"
	"net/http"
)

func Run() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
