package authentication

import (
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/nv4re/go-goo/entity/errors"
)

// The reason that I can't separate authentication, role, permission into each entity because the depend on each other
type User struct {
	Hash      []byte    `json:"hash"`
	Roles     []Role    `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	Info
}

type Info struct {
	Username    string `json:"username"`
	ProfileName string `json:"profile_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

func NewUser(username, password, profileName, email, phone string) (*User, error) {
	u := &User{
		CreatedAt: time.Now(),
		Info: Info{
			Username:    username,
			ProfileName: profileName,
			Email:       email,
			Phone:       phone,
		},
	}

	if err := u.SetPassword(password); err != nil {
		return nil, err
	}

	if err := u.CheckUsername(username); err != nil {
		return nil, err
	}

	if err := u.CheckEmail(email); err != nil {
		return nil, err
	}

	if err := u.CheckPhone(phone); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) SetPassword(password string) error {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Hash = h
	return nil
}

func (u *User) SetInfo(info *Info) error {
	if info == nil {
		return errors.BadRequest
	}

	if info.ProfileName != "" {
		u.ProfileName = info.ProfileName
	}

	if info.Phone != "" {
		u.Phone = info.Phone
	}

	if info.Email != "" {
		u.Email = info.Email
	}

	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Hash, []byte(password))
	return err == nil
}

func (u *User) CheckPermission(p Permission) error {
	if len(u.Roles) == 0 {
		return errors.InsufficientPrivileges
	}

	for _, r := range u.Roles {
		if err := r.CheckPermission(p); err != nil {
			return err
		}
	}
	return nil
}

func (i *Info) CheckUsername(username string) error {
	m, err := regexp.Match("(?i)^[a-z0-9\\-]{6,32}$", []byte(username))
	if err != nil || !m {
		return errors.InvalidUsername
	}
	return nil
}

func (i *Info) CheckEmail(email string) error {
	m, err := regexp.Match("(?i)^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$", []byte(email))
	if err != nil || !m {
		return errors.InvalidEmail
	}
	return nil
}

func (i *Info) CheckPhone(phone string) error {
	m, err := regexp.Match("^[0-9]{9,10}$", []byte(phone))
	if err != nil || !m {
		return errors.InvalidPhoneNumber
	}
	return nil
}
