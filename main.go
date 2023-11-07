package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/items", GetItemsList)
	router.GET("/items/:id", GetItem)
	router.POST("/items", CreateItem)
	router.PATCH("/items/:id", UpdateItem)
	router.PUT("/items/:id", UpdateItem)
	router.DELETE("/items/:id", DeleteItem)

	log.Fatal(http.ListenAndServe(":3000", router))
}
