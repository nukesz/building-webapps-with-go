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
	ID      string `json:"id"`
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
	log.Println("Connected to Postgres!")

	// Migrate the schema
	db.AutoMigrate(&Book{})

	// Create
	db.Create(&Book{ID: "55", Title: "L1212", Author: "Norbert"})
	db.Create(&Book{ID: "66", Title: "S2323", Author: "Unknown"})

	log.Println("Listening...")
	http.Handle("/", showBooks(db))
	http.ListenAndServe(":8080", nil)
}

func showBooks(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read
		var book Book
		db.First(&book, r.FormValue("id"))
		js, err := json.MarshalIndent(book, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}
