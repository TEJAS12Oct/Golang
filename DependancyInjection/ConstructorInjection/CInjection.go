package main

import (
	"fmt"
)

type UserRepository interface {
	GetUserFromDB() string
}

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) GetUserFromDB() string {
	return "User from Database"
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// ****************************************************************************** //

type UserService interface {
	GetUser() string
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}

}

func (s *UserServiceImpl) GetUser() string {
	data := s.repo.GetUserFromDB()
	return "Service Layer → " + data
}

// ****************************************************************************** //

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Handle() {
	result := h.service.GetUser()
	fmt.Println("Constructor Injection -> Handler Layer →", result)

}

func main() {
	repo := NewUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)
	handler.Handle()

}
