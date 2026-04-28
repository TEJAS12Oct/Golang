package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// 🔹 Struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 🔹 In-memory data
var users = []User{
	{ID: 1, Name: "Tejas"},
	{ID: 2, Name: "Rahul"},
}

// 🔹 Helper: JSON Response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// 🔹 Helper: Error Response
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// 🔹 Helper: Parse ID
func parseID(r *http.Request) (int, error) {
	idParam := r.URL.Query().Get("id")
	return strconv.Atoi(idParam)
}

// 🔹 GET: All Users
func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	respondJSON(w, http.StatusOK, users)
}

// 🔹 GET: User by ID
func getUserByID(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	for _, user := range users {
		if user.ID == id {
			respondJSON(w, http.StatusOK, user)
			return
		}
	}

	respondError(w, http.StatusNotFound, "User not found")
}

// 🔹 POST: Create User
func createUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if newUser.Name == "" {
		respondError(w, http.StatusBadRequest, "Name is required")
		return
	}

	for _, u := range users {
		if u.ID == newUser.ID {
			respondError(w, http.StatusBadRequest, "ID already exists")
			return
		}
	}

	users = append(users, newUser)
	respondJSON(w, http.StatusCreated, newUser)
}

// 🔹 PUT: Update User
func updateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		respondError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if updatedUser.Name == "" {
		respondError(w, http.StatusBadRequest, "Name is required")
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			respondJSON(w, http.StatusOK, users[i])
			return
		}
	}

	respondError(w, http.StatusNotFound, "User not found")
}

// 🔹 DELETE: Delete User
func deleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	respondError(w, http.StatusNotFound, "User not found")
}

// 🔹 MAIN
func main() {
	/*
		// GetAllUsers: http://localhost:8080/users
		http.HandleFunc("/users", getUsers)
		// GetUserByID: http://localhost:8080/users/GetUser?id=1
		http.HandleFunc("/users/GetUser", getUserByID)
		// CreateUser: http://localhost:8080/users/CreateUser (POST)
		http.HandleFunc("/users/CreateUser", createUser)
		// UpdateUser: http://localhost:8080/users/UpdateUser?id=1 (PUT)
		http.HandleFunc("/users/UpdateUser", updateUser)
		// DeleteUser: http://localhost:8080/users/DeleteUser?id=1 (DELETE)
		http.HandleFunc("/users/DeleteUser", deleteUser)

		fmt.Println("Server running on port 8080...")
		http.ListenAndServe(":8080", nil)
	*/

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id != "" {
				getUserByID(w, r)
				return //  VERY IMPORTANT
			}
			getUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		case http.MethodPut:
			updateUser(w, r)
		case http.MethodDelete:
			deleteUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
