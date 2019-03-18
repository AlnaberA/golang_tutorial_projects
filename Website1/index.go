package main

import (
	"fmt"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world wide localhost!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
