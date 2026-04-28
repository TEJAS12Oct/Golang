package Handlers

import (
	models "RestAPIDatabaseConnection/Models"
	repository "RestAPIDatabaseConnection/Repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	students, err := repository.GetUsers()
	if err != nil {
		log.Println("ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var student struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&student) // Decode the JSON body into the student struct
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if student.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	err = repository.InsertUser(student.Name)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User created successfully"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if student.ID == 0 || student.Name == "" {
		http.Error(w, "id and name required", http.StatusBadRequest)
		return
	}

	err = repository.UpdateUser(student.ID, student.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User updated successfully"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = repository.DeleteUser(id)
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User deleted successfully"))
}
