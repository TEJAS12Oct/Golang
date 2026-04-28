package main

import (
	db "RestAPIDatabaseConnection/DB"
	handlers "RestAPIDatabaseConnection/Handlers"
	"log"
	"net/http"
)

func main() {
	db.Connect()

	http.HandleFunc("/students/select", handlers.GetUsersHandler)
	http.HandleFunc("/students/create", handlers.CreateUser)
	http.HandleFunc("/students/update", handlers.UpdateUser)
	http.HandleFunc("/students/delete", handlers.DeleteUser)

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
