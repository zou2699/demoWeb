package main

import (
	"io"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, kubernetes!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
