package handler

import "WireDI/Project/service"

// Conntains UserHandler struct which depends on UserService
// and provide Handle() method to print the user details
type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// ✅ ADD THIS METHOD
func (h *UserHandler) Handle() {
	println(h.service.GetUser())
}
