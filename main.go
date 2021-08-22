package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

var port string = "8080"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Listening on localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
