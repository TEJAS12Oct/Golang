// service/user_service.go
package service

import (
	"RESTAPIPOSTGIN/model"
	"RESTAPIPOSTGIN/repository"
	"fmt"
)

type UserService struct { // UserService struct has a field repo of type repository.UserRepo,
	// which is an interface that defines the methods for accessing user data from the database
	repo repository.UserRepo // it is used to interact with the user repository
	// to perform operations related to users
}

func NewUserService(repo repository.UserRepo) *UserService { // NewUserService is a constructor function
	// that takes a UserRepo as an argument and returns a pointer to a UserService
	return &UserService{repo: repo} // it initializes the repo field of the UserService struct with the provided UserRepo
}

func (s *UserService) GetUsers() (interface{}, error) { // GetUsers is a method of the UserService struct that returns
	//  a slice of users and an error
	return s.repo.GetUsers() // it calls the GetUsers method of the UserRepo
	// interface to retrieve the users from the database
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	return s.repo.CreateUser(user)
}

// CreateUser is a method of the UserService struct that takes a user as an argument and returns the created user and an error
// it calls the CreateUser method of the UserRepo interface to create a new user in the database

func (s *UserService) UpdateUser(user model.User) (model.User, error) {

	fmt.Println("Service received:", user)
	return s.repo.UpdateUser(user)

}

func (s *UserService) DeleteUser(user model.User) (model.User, error) {

	fmt.Println("Service received:", user)
	return s.repo.DeleteUser(user.ID)

}
