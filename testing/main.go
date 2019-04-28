package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello %s!", name)
	case "POST":
		fmt.Fprint(w, "Processing POST request")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}
