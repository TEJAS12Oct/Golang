//go:build wireinject
// +build wireinject

package main

import (
	dbconfig "RESTAPIPOSTGIN/dbConfig"
	"RESTAPIPOSTGIN/handler"
	"RESTAPIPOSTGIN/repository"
	"RESTAPIPOSTGIN/service"

	"github.com/google/wire"
)

func InitializeHandler() (*handler.UserHandler, error) {
	wire.Build(
		dbconfig.NewDB,
		repository.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,
	)
	return &handler.UserHandler{}, nil
}
