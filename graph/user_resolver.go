package graph

import (
	"context"
	"twitter"
)

func mapUser(u twitter.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (r *queryResolver) Me(c context.Context) (*User, error) {
	userId, err := twitter.GetUserIdFromContext(c)
	if err != nil {
		return nil, twitter.ErrUnAuthenticate
	}

	return mapUser(twitter.User{
		ID: userId,
	}), nil
}
