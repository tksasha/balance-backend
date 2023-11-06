package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Item struct {
	ID 		int			`json:"id"`
	Name 	string	`json:"name"`
}

func GetItemsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Items List")
}

func GetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode("Unprocessable Entity")

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

	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Item Created")
}

func UpdateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Item Updated")
}

func DeleteItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Item Deleted")
}

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

/*

TODO: add routes for:
[+]    GET /items
[+]    GET /items/:id
[+]   POST /items
[+]  PATCH /items/:id
[+]    PUT /items/:id
[+] DELETE /items/:id

*/
