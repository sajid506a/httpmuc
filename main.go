package main

import (
	"fmt"
	"net/http"
)

func main() {

	httpServer := http.NewServeMux()

	httpServer.HandleFunc("/", Home)

	http.ListenAndServe(":8080", httpServer)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello home")
}
