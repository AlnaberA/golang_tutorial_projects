package main

import (
	"fmt"
	"net/http"
)

//Handles the main page.
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world wide localhost!")
}

func main() {
	//The first param is the route and the second param is the page controller
	http.HandleFunc("/", indexHandler)
	//nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.
	//The second param of ListenAndServe is server configurations
	http.ListenAndServe(":8000", nil)
}
