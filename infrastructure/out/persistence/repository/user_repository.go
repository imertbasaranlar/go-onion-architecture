package repository

import (
	"context"

	"github.com/imertbasaranlar/go-onion-architecture/application/repository"
	"github.com/imertbasaranlar/go-onion-architecture/domain/user"
)

type UserRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (uc *UserRepository) GetByID(c context.Context, id string) (*user.User, error) {
	return &user.User{}, nil
}

func (uc *UserRepository) Save(c context.Context, user *user.User) error {
	return nil
}
