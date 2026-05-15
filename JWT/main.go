package main

import (
	handlers "JWT/Handlers"
	middleware "JWT/Middleware"
	DB "JWT/db"
	"log"
	"net/http"
)

func main() {
	DB.Connect()

	// Public APIs
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/delete", middleware.AuthMiddleware(handlers.DeleteUser))

	// Protected API
	http.HandleFunc("/dashboard", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ Authorized Access"))
	}))

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
