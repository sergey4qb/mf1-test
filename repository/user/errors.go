package user

import "errors"

var (
	errCreateUserFile = errors.New("failed to create user file")
	errUserNotFound   = errors.New("user not found")
)
