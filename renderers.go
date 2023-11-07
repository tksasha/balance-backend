package main

import (
	"encoding/json"
	"net/http"
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
