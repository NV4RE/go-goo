package repository

import (
	"context"

	"github.com/nv4re/go-goo/entity/authentication"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *authentication.User) error
	UpdateUser(ctx context.Context, user *authentication.User) error
	GetUserByUsername(ctx context.Context, username string) (*authentication.User, error)
	SearchUser(ctx context.Context, user string, from, to int) ([]*authentication.User, error)

	CreateRole(ctx context.Context, role *authentication.Role) error
	UpdateRole(ctx context.Context, role *authentication.Role) error
	DeleteRoleByName(ctx context.Context, name string) error
	GetRoleByName(ctx context.Context, name string) (*authentication.Role, error)
	ListRole(ctx context.Context, from, to int) ([]*authentication.Role, error)

	IsHealthy(ctx context.Context) error
}
