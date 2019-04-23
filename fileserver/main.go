package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", generateMarkdownHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":"+port, nil)
}

func generateMarkdownHandler(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
