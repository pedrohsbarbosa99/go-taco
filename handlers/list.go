package handlers

import (
	"encoding/json"
	"example/go-taco/parse"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	f, err := parse.OpenFile("taco.json")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	data, err := parse.GetFoodData(f)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Foods)
}
