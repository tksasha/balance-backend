package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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

	item := Item{}

	OK(w, item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	OK(w, "Item Updated")
}

func DeleteItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	OK(w, "Item Deleted")
}
