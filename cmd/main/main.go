package main

import (
	"github.com/gorilla/mux"
	"github.com/hendra61298/first-project/src/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
