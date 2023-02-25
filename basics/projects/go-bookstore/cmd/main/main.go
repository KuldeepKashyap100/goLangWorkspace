package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuldeep/go-bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Starting server at: 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
