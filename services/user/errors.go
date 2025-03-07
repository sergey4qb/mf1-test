package user

import "errors"

var (
	errInvalidName        = errors.New("name cannot be empty")
	errInvalidEmail       = errors.New("email cannot be empty")
	errInvalidFormatEmail = errors.New("invalid format email")
)
