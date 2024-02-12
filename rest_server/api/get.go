package api

import (
	"encoding/json"
	"net/http"
	"rest_server/data"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"] // if an id is sent as query param we will send 1, otherwise all
	if id != nil {
		finalId, err := strconv.Atoi(id[0]) // string to int conversion, always is an array of params
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "invalid", http.StatusBadRequest)
		}
	} else {

		json.NewEncoder(w).Encode(data.GetAll()) // using the encoder instead of marshal or unmarshal to handle the stream
	}

}
