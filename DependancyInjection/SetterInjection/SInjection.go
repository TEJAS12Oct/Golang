package main

import "fmt"

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
	SetRepository(repo UserRepository) // 👈 Setter method
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) SetRepository(repo UserRepository) {
	s.repo = repo
}

func (s *UserServiceImpl) GetUser() string {

	if s.repo == nil {
		return "Repository not initialized!"
	}

	data := s.repo.GetUserFromDB()
	return "Service Layer → " + data
}

// ****************************************************************************** //

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Handle() {
	result := h.service.GetUser()
	fmt.Println("Setter Injection -> Handler Layer →", result)
}

// ****************************************************************************** //

func main() {

	repo := NewUserRepository()

	service := NewUserService()

	service.SetRepository(repo)

	handler := NewUserHandler(service)

	handler.Handle()
}
