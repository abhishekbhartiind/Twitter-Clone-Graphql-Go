package twitter

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrValidation         = errors.New("validation error")
	ErrInvalidAccessToken = errors.New("invalid access")
	ErrNoUserInContext    = errors.New("no user id in context")
	ErrGenTokenAccess     = errors.New("generate access token error")
	ErrUnAuthenticate     = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
)
