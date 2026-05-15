package main

import "fmt"

type UserRepository interface {
	GetUserFromDB() string
}

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) GetUserFromDB() string {
	return "User from Database"
}

type UserService struct{}

func (s *UserService) GetUser(repo UserRepository) string {
	data := repo.GetUserFromDB()
	return "Service Layer → " + data
}

func main() {

	repo := &UserRepositoryImpl{}
	service := &UserService{}

	result := service.GetUser(repo)

	fmt.Println(result)
}
