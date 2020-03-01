package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Item is a struct that groups all necessary fields into a single unit
type Item struct {
	Data      string `json:"data"`
	OtherData int    `json:"otherData"`
}

var data []Item = []Item{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/add", addItem).Methods("POST")

	http.ListenAndServe(":5000", router)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)
	data = append(data, newItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
