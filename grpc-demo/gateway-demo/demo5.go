package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("localhost/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "localhost", r.URL.EscapedPath())
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
