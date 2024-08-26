package main

import (
	"fmt"
	"net/http"

	"github.com/Mpablom/candyblogback/api"
)

func main() {
	http.HandleFunc("/", api.HelloHandler)
	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
