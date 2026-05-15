// handler/user_handler.go
package handler

import (
	"RESTAPIPOSTGIN/model"
	"RESTAPIPOSTGIN/service"
	"fmt"
	"strconv"

	//Module name is RESTAPIPOSTGIN, so import path is RESTAPIPOSTGIN/service ,
	// if module name is different then import path will be different	,
	// Module name is defined in go.mod file
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct { // UserHandler struct has a field service of type *service.UserService
	service *service.UserService // UserService is defined in service package,
	// and it has a method GetUsers which returns a slice of users and an error
}

func NewUserHandler(service *service.UserService) *UserHandler { // NewUserHandler is a constructor function
	// that takes a UserService as an argument and returns a pointer to a UserHandler
	return &UserHandler{service: service} // it initializes the service field of the UserHandler struct with the provided UserService
}

func (h *UserHandler) GetUsers(c *gin.Context) { // GetUsers is a method of UserHandler struct that takes a gin.Context as an argument
	users, err := h.service.GetUsers() // it calls the GetUsers method of the UserService
	// and assigns the returned slice of users to the variable users and the error to the variable err
	if err != nil { // if there is an error, it returns a JSON response with the error message and
		// a status code of 500 (Internal Server Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //	 gin.H is a shortcut for map[string]interface{}
		// and it is used to create a JSON response with the error message
		return // if there is no error, it returns a JSON response with
		// the slice of users and a status code of 200 (OK)
	}
	c.JSON(http.StatusOK, users) // it returns a JSON response with the slice of users and a status code of 200 (OK)
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User

	// 🔹 Bind JSON → struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	// 🔹 Call service
	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// 🔹 Response
	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	// 🔹 Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	var user model.User

	// 🔹 Bind JSON → struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	user.ID = id
	updatedUser, err := h.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}
	fmt.Println("ID from URL:", id)
	fmt.Println("User after bind:", user)
	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// 🔹 Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	var user model.User

	// 🔹 Bind JSON → struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	user.ID = id
	updatedUser, err := h.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}
	fmt.Println("ID from URL:", id)
	fmt.Println("User after bind:", user)
	c.JSON(http.StatusOK, updatedUser)
}
