package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "user", request.URL.EscapedPath())
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
