package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Ignored int    `json:"-"`
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=webapp-demo user=admin password=admin sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to Postgres %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&Book{})

	log.Println("Connected to Postgres!")
	log.Println("Listening...")
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
