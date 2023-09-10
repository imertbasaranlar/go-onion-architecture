package controller

import (
	"context"

	"github.com/imertbasaranlar/go-onion-architecture/application/usecase"
	"github.com/imertbasaranlar/go-onion-architecture/infrastructure/out/persistence/repository"
)

func getByID() {
	userRepository := repository.NewUserRepository()
	userGetByIdUseCase := usecase.NewUserGetByIdUseCase(userRepository)

	user, _ := userGetByIdUseCase.Run(context.Background())
	if user != nil {
	}
}
