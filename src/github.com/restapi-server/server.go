package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	port := 9083
	bindRoutes(r)
	fmt.Println(fmt.Sprintf("api server is running on %d", port))
	if e := http.ListenAndServe(fmt.Sprintf(":%d", port), r); e != nil {
		log.Print(e)
	}
}
