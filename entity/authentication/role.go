package authentication

import (
	"regexp"
	"time"

	"github.com/nv4re/go-goo/entity/errors"
)

type Role struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permission"`
	CreatedAt   time.Time    `json:"created_at"`
}

func NewRole(name, description string, permissions []Permission) *Role {
	return &Role{
		name,
		description,
		permissions,
		time.Now(),
	}
}

func (r *Role) CheckName(name string) error {
	m, err := regexp.Match("^[A-Z0-9_]{4,32}$", []byte(name))
	if err != nil || !m {
		return errors.InvalidUsername
	}
	return nil
}

func (r *Role) CheckPermission(p Permission) error {
	for _, pm := range r.Permissions {
		if p == pm {
			return nil
		}
	}
	return errors.InsufficientPrivileges
}
