package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var DB *sql.DB

func main() {
	DB = OpenDBConnection()

	defer DB.Close()

	router := httprouter.New()

	router.GET("/items", GetItemsList)
	router.GET("/items/:id", GetItem)
	router.POST("/items", CreateItem)
	router.PATCH("/items/:id", UpdateItem)
	router.PUT("/items/:id", UpdateItem)
	router.DELETE("/items/:id", DeleteItem)

	log.Fatal(http.ListenAndServe(":3000", router))
}
