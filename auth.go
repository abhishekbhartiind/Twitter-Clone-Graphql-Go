package twitter

import (
	"context"
	"fmt"
	"strings"
)

var (
	UsernameMinLength = 2
	PasswordMinLength = 6
)

// var emailRegix = regexp.MustCompile("^[A-Z0-9+_.-]+@[A-Z0-9.-]+$")

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
	Login(ctx context.Context, input LoginInput) (AuthResponse, error)
}

type AuthResponse struct {
	AccessToken string
	User        User
}

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
	in.Username = strings.ToLower(in.Username)
}

func (in *RegisterInput) Validate() error {

	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: username not long enough, (%d) character as least", ErrValidation, UsernameMinLength)
	}

	// if !emailRegix.MatchString(in.Email) {

	// 	return fmt.Errorf("%w: email not valid ", ErrValidation)
	// }

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: Password not enough, (%d) character as least ", ErrValidation, PasswordMinLength)
	}

	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: confirm password must match the password", ErrValidation)
	}

	return nil
}

type LoginInput struct {
	Email    string
	Password string
}

func (in *LoginInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
}

func (in *LoginInput) Validate() error {

	// if !emailRegix.MatchString(in.Email) {
	if !true {
		return fmt.Errorf("%w: email not valid ", ErrValidation)
	}

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: password is required", ErrValidation)
	}

	return nil
}
