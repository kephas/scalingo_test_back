package main

//go:generate bower install

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
)

func HumanReadableBytes(ibytes int) string {
	var bytes = float64(ibytes)
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
		return fmt.Sprintf("%.0f %s", (bytes / math.Pow(1024, power)), suffix)
	} else {
		return fmt.Sprintf("%.2f %s", (bytes / math.Pow(1024, power)), suffix)
	}
}

func StaticFileHandler(name string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, name)
	}
}

var searchLimit int = 100

func SearchPage(w http.ResponseWriter, r *http.Request) {
	search, err := SearchGithub(r.URL.Query().Get("query"), searchLimit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	CreateWorkers()
	for index := range search {
		DispatchRepo(&search[index])
	}
	CloseChannelsAndWait()

	tmpl, err := template.New("search.html").Funcs(template.FuncMap{"human": HumanReadableBytes}).ParseFiles("search.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, search)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", StaticFileHandler("home.html"))
	r.HandleFunc("/bootstrap.css", StaticFileHandler("bower_components/bootstrap/dist/css/bootstrap.css"))
	r.HandleFunc("/search", SearchPage)

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "port, p", Value: "8000", EnvVar: "PORT"},
		cli.IntFlag{Name: "limit, l", Value: 100, Destination: &searchLimit},
		cli.IntFlag{Name: "workers, w", Value: 10, Destination: &WorkerPoolSize},
	}
	app.Action = func(c *cli.Context) error {
		log.Fatal(http.ListenAndServe(":" + c.String("port"), r))
		return nil
	}

	app.Run(os.Args)
}
