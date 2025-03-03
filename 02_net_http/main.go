package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	domain "github.com/xavicci/FULLAPIRESTGO/domain"
)

func main() {
	http.HandleFunc("create-user", createUserHandler)
	http.HandleFunc("get-user", getUserHandler)
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]
	fmt.Fprintf(w, "Request user id: %s", id)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if user.Age < 18 {
		http.Error(w, "User must be 18 years", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "USer create successfully")
}
