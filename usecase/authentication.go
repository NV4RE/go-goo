package usecase

import (
	"context"

	"github.com/nv4re/go-goo/entity/authentication"
	"github.com/nv4re/go-goo/entity/errors"
	"github.com/nv4re/go-goo/repository"
)

type authenticationUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthenticationUseCase(authRepo repository.AuthRepository) AuthenticationUseCase {
	return &authenticationUseCase{authRepo}
}

func (a *authenticationUseCase) Register(ctx context.Context, username, password, profileName, email, phone string) error {
	u, err := authentication.NewUser(username, password, profileName, email, phone)
	if err != nil {
		return err
	}

	err = a.authRepo.CreateUser(ctx, u)
	return err
}

func (a *authenticationUseCase) Login(ctx context.Context, username, password string) (string, error) {
	u, err := a.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	m := u.CheckPassword(password)
	if m != true {
		return "", errors.PasswordNotMatch
	}

	return username, nil
}

func (a *authenticationUseCase) ChangePassword(ctx context.Context, user *authentication.User, password, oldPassword string) error {
	m := user.CheckPassword(oldPassword)
	if m != true {
		return errors.PasswordNotMatch
	}

	err := user.SetPassword(password)
	if err != nil {
		return err
	}

	err = a.authRepo.UpdateUser(ctx, user)
	return err
}

func (a *authenticationUseCase) ChangeInfo(ctx context.Context, user *authentication.User, username string, info *authentication.Info) error {
	err := user.CheckPermission(authentication.PermissionUpdateUser)
	if err != nil {
		return err
	}

	u, err := a.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	err = u.SetInfo(info)
	if err != nil {
		return err
	}

	err = a.authRepo.UpdateUser(ctx, user)
	return err
}

func (a *authenticationUseCase) ChangeMyInfo(ctx context.Context, user *authentication.User, info *authentication.Info) error {
	err := user.SetInfo(info)
	if err != nil {
		return err
	}

	err = a.authRepo.UpdateUser(ctx, user)
	return err
}

func (a *authenticationUseCase) GetInfo(ctx context.Context, user *authentication.User, username string) (*authentication.Info, error) {
	err := user.CheckPermission(authentication.PermissionViewUser)
	if err != nil {
		return nil, err
	}

	u, err := a.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &u.Info, nil
}

func (a *authenticationUseCase) GetMyInfo(_ context.Context, user *authentication.User) (*authentication.Info, error) {
	return &user.Info, nil
}

func (a *authenticationUseCase) SearchUser(ctx context.Context, user *authentication.User, username string, from, to int) ([]*authentication.Info, int, error) {
	err := user.CheckPermission(authentication.PermissionViewUser)
	if err != nil {
		return nil, 0, err
	}

	users, err := a.authRepo.SearchUser(ctx, username, from, to)
	if err != nil {
		return nil, 0, err
	}

	infos := make([]*authentication.Info, len(users))
	for i, u := range users {
		infos[i] = &u.Info
	}

	return infos, 0, nil
}

func (a *authenticationUseCase) AssignRole(ctx context.Context, user *authentication.User, username, roleName string) error {
	err := user.CheckPermission(authentication.PermissionAssignRole)
	if err != nil {
		return err
	}

	u, err := a.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	r, err := a.authRepo.GetRoleByName(ctx, roleName)
	if err != nil {
		return err
	}

	u.Roles = append(u.Roles, *r)

	return a.authRepo.UpdateUser(ctx, u)
}

func (a *authenticationUseCase) CreateRole(ctx context.Context, user *authentication.User, name, description string, permissions []authentication.Permission) (*authentication.Role, error) {
	err := user.CheckPermission(authentication.PermissionCreateRole)
	if err != nil {
		return nil, err
	}

	r := authentication.NewRole(name, description, permissions)
	err = a.authRepo.CreateRole(ctx, r)
	return r, err
}

func (a *authenticationUseCase) UpdateRole(ctx context.Context, user *authentication.User, name string, permissions []authentication.Permission) (*authentication.Role, error) {
	err := user.CheckPermission(authentication.PermissionUpdateRole)
	if err != nil {
		return nil, err
	}

	r, err := a.authRepo.GetRoleByName(ctx, name)
	if err != nil {
		return nil, err
	}

	r.Permissions = permissions
	err = a.authRepo.UpdateRole(ctx, r)
	return r, err
}

func (a *authenticationUseCase) DeleteRole(ctx context.Context, user *authentication.User, name string) error {
	err := user.CheckPermission(authentication.PermissionDeleteRole)
	if err != nil {
		return err
	}

	return a.authRepo.DeleteRoleByName(ctx, name)
}

func (a *authenticationUseCase) GetRole(ctx context.Context, user *authentication.User, name string) (*authentication.Role, error) {
	err := user.CheckPermission(authentication.PermissionViewRole)
	if err != nil {
		return nil, err
	}

	return a.authRepo.GetRoleByName(ctx, name)
}

func (a *authenticationUseCase) ListRole(ctx context.Context, user *authentication.User, from, to int) ([]*authentication.Role, int, error) {
	err := user.CheckPermission(authentication.PermissionViewRole)
	if err != nil {
		return nil, 0, err
	}

	roles, err := a.authRepo.ListRole(ctx, from, to)
	if err != nil {
		return nil, 0, err
	}
	return roles, 0, err
}
