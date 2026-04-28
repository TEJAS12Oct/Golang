package Handlers

import (
	"JWT/Repository"
	models "JWT/models"
	"JWT/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" || user.Password == "" {
		http.Error(w, "username & password required", 400)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	err := Repository.CreateUser(user)
	if err != nil {
		http.Error(w, "User already exists", 400)
		return
	}

	w.Write([]byte("User registered successfully"))
}

// 🔹 LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)

	user, err := Repository.GetUserByUsername(creds.Username)
	if err != nil {
		http.Error(w, "User not found", 401)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid password", 401)
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Username)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
