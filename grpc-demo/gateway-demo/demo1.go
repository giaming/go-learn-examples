package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "/a", request.URL.EscapedPath())
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "default", request.URL.EscapedPath())
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
