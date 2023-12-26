package api

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
