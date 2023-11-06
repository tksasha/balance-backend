package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func JSON(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(i)
}

func OK(w http.ResponseWriter, i interface{}) {
	JSON(w, http.StatusOK, i)
}

func Created(w http.ResponseWriter, i interface{}) {
	JSON(w, http.StatusCreated, i)
}

func UnprocessableEntity(w http.ResponseWriter, i interface{}) {
	JSON(w, http.StatusUnprocessableEntity, i)
}

func NotFound(w http.ResponseWriter, i interface{}) {
	JSON(w, http.StatusNotFound, i)
}

type Item struct {
	ID 		int			`json:"id"`
	Name 	string	`json:"name"`
}

func GetItemsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	OK(w, "Items List")
}

func GetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		UnprocessableEntity(w, "Unprocessable Entity")

		return
	}

	if id <= 0 {
		NotFound(w, "Not Found")

		return
	}

	item := Item{id, "Pretty Red Dress"}

	OK(w, item)
}

func CreateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Created(w, "Item Created")
}

func UpdateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	OK(w, "Item Updated")
}

func DeleteItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	OK(w, "Item Deleted")
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
