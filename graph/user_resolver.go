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

func (r *Resolver) Me(c context.Context) (*User, error) {
	panic("implement me")
}
