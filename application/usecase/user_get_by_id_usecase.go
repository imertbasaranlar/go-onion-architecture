package usecase

import (
	"context"

	"github.com/imertbasaranlar/go-onion-architecture/application/repository"

	"github.com/imertbasaranlar/go-onion-architecture/domain/user"
)

type (
	IUserGetByIdUseCase interface {
		Run(ctx context.Context) (*user.User, error)
	}
	UserGetByIdUseCase struct {
		userRepository repository.UserRepository
	}
)

func NewUserGetByIdUseCase(userRepository repository.UserRepository) IUserGetByIdUseCase {
	return &UserGetByIdUseCase{
		userRepository: userRepository,
	}
}

func (uc *UserGetByIdUseCase) Run(ctx context.Context) (*user.User, error) {
	return &user.User{}, nil
}
