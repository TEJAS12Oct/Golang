package service

import "WireDI/Project/repository"

// 👉 Service depends on repository
type UserService struct {
	repo *repository.UserRepo
}

// Constructor for UserService which takes UserRepo as a parameter and return a Pointer to UserService	]
func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

// ✅ ADD THIS METHOD
func (s *UserService) GetUser() string {
	return s.repo.GetUser() // calls GetUser() from repository and returns the result
}
