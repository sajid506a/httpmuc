package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type User struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Age         string `json:"age"`
}

var userCache = make(map[int]User)
var cacheMutex sync.RWMutex
var users []User

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("POST /users", createUser)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func Home(w http.ResponseWriter, r *http.Request) {
	usersString := strings.Join(usersToStringSlice(users), "\n")

	fmt.Fprintf(w, usersString)
}

func usersToStringSlice(users []User) []string {
	result := make([]string, len(users))
	for i, u := range users {
		result[i] = fmt.Sprintf("%+v", u)
	}
	return result
}
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()
	users = append(users, user)
	w.WriteHeader(http.StatusNoContent)
}
