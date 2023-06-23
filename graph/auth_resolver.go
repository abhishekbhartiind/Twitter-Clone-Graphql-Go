package graph

import (
	"context"
	"twitter"
)

func mapAuthResponse(a twitter.AuthResponse) *AuthResponse {
	return &AuthResponse{
		AccessToken: a.AccessToken,
		// User:,
	}
}

func (m *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {

	_, err := m.Resolver.AuthService.Register(
		ctx, twitter.RegisterInput{
			Username:        input.Username,
			Email:           input.Email,
			Password:        input.Password,
			ConfirmPassword: input.Password,
		},
	)

	if err != nil {

	}

	panic("implement me")
}
func (m *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	panic("implement me")
}
