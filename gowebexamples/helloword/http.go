package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "Hello, you've requested:%s \n ", request.URL.Path)
	})

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
