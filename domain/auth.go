package domain

import (
	"context"
	"errors"
	"fmt"
	"log"
	"twitter"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo twitter.UserRepo
}

func NewAuthService(ur twitter.UserRepo) *AuthService {
	return &AuthService{
		UserRepo: ur,
	}
}

func (as *AuthService) Register(c context.Context, input twitter.RegisterInput) (twitter.AuthResponse, error) {

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return twitter.AuthResponse{}, err
	}

	if _, err := as.UserRepo.GetByUsername(c, input.Username); errors.Is(err, twitter.ErrNotFound) {
		return twitter.AuthResponse{}, twitter.ErrUsernameTaken
	}

	if _, err := as.UserRepo.GetByEmail(c, input.Email); errors.Is(err, twitter.ErrNotFound) {
		return twitter.AuthResponse{}, twitter.ErrEmailTaken
	}

	user := twitter.User{
		Username: input.Username,
		Email:    input.Email,
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 6)
	if err != nil {
		return twitter.AuthResponse{}, fmt.Errorf("error hashing password: %v", err)
	}

	user.Password = string(hashPassword)

	log.Println("just testing user", user)

	createUser, err := as.UserRepo.Create(c, user)
	if err != nil {
		return twitter.AuthResponse{}, fmt.Errorf("error while creating users:%v", err)
	}

	return twitter.AuthResponse{
		AccessToken: "access token",
		User:        createUser,
	}, nil
}

func (as *AuthService) Login(c context.Context, input twitter.LoginInput) (twitter.AuthResponse, error) {

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return twitter.AuthResponse{}, err
	}

	user, err := as.UserRepo.GetByEmail(c, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, twitter.ErrNotFound):
			return twitter.AuthResponse{}, twitter.ErrCredentials
		default:
			return twitter.AuthResponse{}, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return twitter.AuthResponse{}, twitter.ErrCredentials
	}

	return twitter.AuthResponse{
		AccessToken: "access token",
		User:        user,
	}, nil
}
