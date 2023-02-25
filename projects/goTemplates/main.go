package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(":8080", nil)
}
