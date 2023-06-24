package graph

import (
	"context"
	"errors"
	"twitter"
)

func mapAuthResponse(a twitter.AuthResponse) *AuthResponse {
	return &AuthResponse{
		AccessToken: a.AccessToken,
		User:        mapUser(a.User),
	}
}

func (m *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {

	res, err := m.AuthService.Register(
		ctx, twitter.RegisterInput{
			Username:        input.Username,
			Email:           input.Email,
			Password:        input.Password,
			ConfirmPassword: input.Password,
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, twitter.ErrValidation) || errors.Is(err, twitter.ErrUsernameTaken):
			return nil, buildBadRequest(ctx, err)

		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}

func (m *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {

	res, err := m.AuthService.Login(ctx, twitter.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		switch {
		case errors.Is(err, twitter.ErrValidation) || errors.Is(err, twitter.ErrNotFound):
			return nil, buildBadRequest(ctx, err)
		default:
			return nil, err
		}
	}
	return mapAuthResponse(res), nil
}
