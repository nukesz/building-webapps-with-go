package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	http.HandleFunc("/", showBooks())
	http.ListenAndServe(":8080", nil)
}

func showBooks() func(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatalf("Could not parse files %v", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
