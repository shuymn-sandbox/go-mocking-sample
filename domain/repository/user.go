package repository

import (
	"context"

	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id int) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}
