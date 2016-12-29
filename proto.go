package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func SearchPage(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage)
	r.HandleFunc("/search", SearchPage)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
