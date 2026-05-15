package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // gin.Default() is a function from the gin package that returns a new
	// instance of the gin engine with the default middleware (logger and recovery) already attached to it.

	userHandler, err := InitializeHandler() // InitializeHandler is a function that initializes
	// 	the user handler and returns it along with any error that occurs during the initialization process
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users/Create", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	r.Run(":8080")
}
