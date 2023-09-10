package repository

import (
	"context"

	"github.com/imertbasaranlar/go-onion-architecture/domain/user"
)

type UserRepository interface {
	Save(c context.Context, user *user.User) error
	GetByID(c context.Context, id string) (*user.User, error)
}
