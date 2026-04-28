package main

import (
	middleware "JWT/Middleware"
	DB "JWT/db"
	handlers "JWT/handlers"
	"log"
	"net/http"
)

func main() {
	DB.Connect()

	// Public APIs
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	// Protected API
	http.HandleFunc("/dashboard", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ Authorized Access"))
	}))

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
