package twitter

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCredentials   = errors.New("email/password wrong combinations")
	ErrEmailTaken    = errors.New("email is taken ")
	ErrUsernameTaken = errors.New("user name is taken ")
)

type UserServices interface {
	GetById(c context.Context, id string) (User, error)
}

type UserRepo interface {
	Create(c context.Context, user User) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetById(c context.Context, id string) (User, error)
	GetByIds(c context.Context, ids []string) ([]User, error)
}

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
