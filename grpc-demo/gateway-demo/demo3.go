package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/image/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "image", request.URL.EscapedPath())
	})
	http.HandleFunc("/image/gif", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "gif", request.URL.EscapedPath())
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
