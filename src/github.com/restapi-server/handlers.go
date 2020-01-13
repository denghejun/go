package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleBooksPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "You request is {title: %s, page: %s} \n", title, page)
}
