//go:build wireinject
// +build wireinject

package main

import (
	"WireDI/Project/handler"
	"WireDI/Project/repository"
	"WireDI/Project/service"

	"github.com/google/wire"
)

func InitializeUserHandler() *handler.UserHandler {
	wire.Build(
		repository.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,
	)
	return &handler.UserHandler{}
}
