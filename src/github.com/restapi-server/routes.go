package main

import "github.com/gorilla/mux"

func bindRoutes(r *mux.Router) *mux.Route {
	return r.HandleFunc("/books/{title}/page/{page}", handleBooksPage).Methods("GET")
}
