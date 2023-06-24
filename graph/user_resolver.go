package graph

import (
	"context"
	"twitter"
)

func mapUser(u twitter.User) *User {
	return &User{
		Email:    u.Email,
		Username: u.Username,
		CreateAt: u.CreateAt,
	}
}

func (r *Resolver) Me(c context.Context) (*User, error) {
	panic("implement me")
}
