package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from server 1")
}
