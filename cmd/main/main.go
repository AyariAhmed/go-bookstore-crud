package main

import (
	"fmt"
	"github.com/AyariAhmed/bookstore-crud/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on port 9010 (http://localhost:9010)")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
