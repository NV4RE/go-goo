package errors

import (
	"errors"
)

// TODO error should not be an entity, because others entity depend on it
var (
	BadRequest = errors.New("bad request")

	InvalidUsername    = errors.New("invalid username")
	InvalidEmail       = errors.New("invalid email")
	InvalidPhoneNumber = errors.New("invalid phone number")
	InvalidRole        = errors.New("invalid role name")
	PasswordNotMatch   = errors.New("password not match")

	NotFoundError = errors.New("not found error")

	InsufficientPrivileges = errors.New("insufficient privileges")

	DeserializationFailed = errors.New("deserialization failed")
	InternalServerError   = errors.New("internal server error")
)
