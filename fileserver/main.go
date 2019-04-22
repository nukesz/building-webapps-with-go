package main

import (
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", generateMarkdownHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8080", nil)
}

func generateMarkdownHandler(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
