package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello, yes I am Bao\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
