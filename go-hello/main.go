package main

import (
	"encoding/json"
	"fmt"
	"go-hello/helpers"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User
var nextID = 1

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var body User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	body.ID = nextID
	nextID++
	users = append(users, body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {

	fmt.Println(helpers.HTTP)
	fmt.Println(helpers.HTTPS)
	fmt.Println(helpers.PORT_80)
	fmt.Println(helpers.PORT_90)

	http.HandleFunc("GET /users", getUsersHandler)
	http.HandleFunc("POST /users/create", createUserHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
