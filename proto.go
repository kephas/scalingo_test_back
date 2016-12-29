package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StaticFileHandler(name string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, name)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func SearchPage(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", StaticFileHandler("home.html"))
	r.HandleFunc("/bootstrap.css", StaticFileHandler("bower_components/bootstrap/dist/css/bootstrap.css"))
	r.HandleFunc("/search", SearchPage)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
