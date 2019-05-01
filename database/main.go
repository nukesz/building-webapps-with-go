package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Ignored int    `json:"-"`
}

func main() {
	http.HandleFunc("/", showBooks)
	http.ListenAndServe(":8080", nil)
}

func showBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz", 5}
	js, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
