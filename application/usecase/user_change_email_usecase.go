package usecase

import (
	"context"

	"github.com/imertbasaranlar/go-onion-architecture/application/repository"
)

type (
	IUserChangeEmailUseCase interface {
		Run(ctx context.Context, id string, email string) error
	}
	UserChangeEmailUseCase struct {
		userRepository repository.UserRepository
	}
)

func NewUserChangeEmailUseCaseUseCase(userRepository repository.UserRepository) IUserChangeEmailUseCase {
	return &UserChangeEmailUseCase{
		userRepository: userRepository,
	}
}

func (uc *UserChangeEmailUseCase) Run(ctx context.Context, id string, email string) error {
	user, _ := uc.userRepository.GetByID(context.Background(), id)
	if user != nil {
		user.ChangeEmail(email)
		return uc.userRepository.Save(context.Background(), user)
	}
	return nil
}
