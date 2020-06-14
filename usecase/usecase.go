package usecase

import (
	"context"

	"github.com/nv4re/go-goo/entity/authentication"
)

type AuthenticationUseCase interface {
	Register(ctx context.Context, username, password, profileName, email, phone string) error
	Login(ctx context.Context, username, password string) (string, error) // Return token
	ChangePassword(ctx context.Context, user *authentication.User, password, oldPassword string) error
	ChangeInfo(ctx context.Context, user *authentication.User, username string, info *authentication.Info) error
	ChangeMyInfo(ctx context.Context, user *authentication.User, info *authentication.Info) error
	GetInfo(ctx context.Context, user *authentication.User, username string) (*authentication.Info, error)
	GetMyInfo(ctx context.Context, user *authentication.User) (*authentication.Info, error)
	SearchUser(ctx context.Context, user *authentication.User, username string, from, to int) ([]*authentication.Info, int, error)
	AssignRole(ctx context.Context, user *authentication.User, username, roleName string) error

	CreateRole(ctx context.Context, user *authentication.User, name, description string, permissions []authentication.Permission) (*authentication.Role, error)
	UpdateRole(ctx context.Context, user *authentication.User, name string, permissions []authentication.Permission) (*authentication.Role, error)
	DeleteRole(ctx context.Context, user *authentication.User, name string) error
	GetRole(ctx context.Context, user *authentication.User, name string) (*authentication.Role, error)
	ListRole(ctx context.Context, user *authentication.User, from, to int) ([]*authentication.Role, int, error)
}

type SystemUseCase interface {
	IsAllHealthy(ctx context.Context) error
}
