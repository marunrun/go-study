package main

import (
	"fmt"
	"net/http"
	"path"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "welcome to my website!")
	})
	_, filename, _, _ := runtime.Caller(0)
	s := path.Dir(filename)

	fs := http.FileServer(http.Dir(s + "/static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	_ = http.ListenAndServe(":8082", nil)
}
