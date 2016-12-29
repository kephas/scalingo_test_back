package main

import (
	"fmt"
	"math"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func HumanReadableBytes(bytes float64) string {
	var magnitude = math.Log2(bytes) / 10
	var power float64
	var suffix string

	if magnitude > 3 {
		power = 3
		suffix = "Gb"
	} else if magnitude > 2 {
		power = 2
		suffix = "Mb"
	} else if magnitude > 1 {
		power = 1
		suffix = "kb"
	} else {
		power = 0
		suffix = "b"
	}

	if power == 0 {
		return fmt.Sprintf("%.0f %s", (bytes/math.Pow(1024, power)), suffix)
	} else {
		return fmt.Sprintf("%.2f %s", (bytes/math.Pow(1024, power)), suffix)
	}
}


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
