package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello home")
}
