package handlers

import (
	"encoding/json"
	"net/http"
)

type dataItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var store = []dataItem{
	{ID: "1", Name: "Item One"},
	{ID: "2", Name: "Item Two"},
	{ID: "3", Name: "Item Three"},
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for _, item := range store {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
}
