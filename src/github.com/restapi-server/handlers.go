package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func handleBooksPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	pageVar := vars["page"]
	i, _ := strconv.Atoi(pageVar) // covert to int.
	book := Book{Title: title, Page: i, description: "this is desc.", Tag: "1", Unknown: "unknown value."}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
	//fmt.Fprintf(w, "You request is {title: %s, page: %s} \n", book.title, book.page)
}
