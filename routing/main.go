package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/", homeHandler)

	r.GET("/posts", postsIndexHandler)
	r.POST("/posts", postsCreateHandler)

	r.GET("/posts/:id", postShowHandler)
	r.PUT("/posts/:id", postUpdateHandler)
	r.GET("/posts/:id/edit", postEditHandler)

	fmt.Println("Server started on :8080")

	http.ListenAndServe(":8080", r)
}

func homeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Home")
}

func postsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "postsIndexHandler")
}

func postsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "postsCreateHandler")
}

func postShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "postShowHandler")
}

func postUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "postUpdateHandler")
}

func postEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "postEditHandler")
}
