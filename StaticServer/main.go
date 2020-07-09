package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// parse command line parameters
	port := flag.String("p", "8000", "port") // return values are pointers
	dir := flag.String("d", ".", "dir")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir))) // use http.FileServer to serve static files
	log.Printf("Serving %s on HTTP port %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
