// Gorilla Mux is REST API framework developed by Golang developer communities which is quite useful for extensive REST API Operations

package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	
	// Declare new Router
	router := mux.NewRouter()

	// GET API Call
	router.HandleFunc("/hello", handler).Methods("GET")

	// Listen the server
	http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
