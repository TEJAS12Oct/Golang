package main

import "fmt"

type UserRepository interface {
	GetUser() string
}

type DBRepository struct{}

func (d *DBRepository) GetUser() string {
	return "User from Database"
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser() string {
	return "Service → " + s.repo.GetUser()
}

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Handle() {
	result := h.service.GetUser()
	fmt.Println("Handler →", result)
}

func main() {
	repo := &DBRepository{}
	service := NewUserService(repo)
	handler := NewUserHandler(service)
	handler.Handle()
}
