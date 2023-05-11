package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello, yes I am Bao\n")
}

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
