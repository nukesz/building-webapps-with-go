package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	adminRouter := mux.NewRouter()
	adminRouter.HandleFunc("/admin	", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the admin page, you must be Groot!")
	})

	// create common middleware to be shared across routes
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	common := negroni.New(
		recovery,
		negroni.NewLogger(),
		negroni.HandlerFunc(myFirstMiddleware),
		negroni.HandlerFunc(mySecondMiddleware),
	)

	// create a new negroni for the hello middleware
	// using the common middleware as a base
	router.PathPrefix("/hello").Handler(common.With(
		negroni.Wrap(router),
	))
	// create a new negroni for the admin middleware
	// using the common middleware as a base
	router.PathPrefix("/admin").Handler(common.With(
		negroni.HandlerFunc(tokenValidatorMiddleware),
		negroni.Wrap(adminRouter),
	))

	http.ListenAndServe(":8080", router)
}

func tokenValidatorMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Query().Get("token") == "" {
		http.Error(rw, "Not Authorized", 401)
	} else {
		// Could be validated..
		next(rw, r)
	}
}

func myFirstMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("logging on the way there...")

	next(rw, r)
	log.Println("logging on the way back...")
}

func mySecondMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//fmt.Fprintf(rw, "Second middleware")
	next(rw, r)
}
