package main

import (
	"encoding/json"
	// "fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func JSON(w io.Writer, i interface{}) {
	json.NewEncoder(w).Encode(i)
}

type Item struct {
	ID 		int			`json:"id"`
	Name 	string	`json:"name"`
}

func GetItemsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, "Items List")
}

func GetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		JSON(w, "Unprocessable Entity")

		return
	}

	if id <= 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode("Not Found")

		return
	}

	item := Item{id, "Pretty Red Dress"}

	w.Header().Set("Content-Type", "application/json")
	JSON(w, item)
}

func CreateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, "Create Item")
}

func UpdateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, "Update Item")
}

func main() {
	router := httprouter.New()

	router.GET("/items", GetItemsList)
	router.GET("/items/:id", GetItem)
	router.POST("/items", CreateItem)
	router.PATCH("/items/:id", UpdateItem)
	router.PUT("/items/:id", UpdateItem)

	log.Fatal(http.ListenAndServe(":3000", router))
}

/*

TODO: add routes for:
[+]    GET /items
[+]    GET /items/:id
[+]   POST /items
[ ]  PATCH /items/:id
[ ]    PUT /items/:id
[ ] DELETE /items/:id
*/
