package postgres

import (
	"context"
	"fmt"
	"twitter"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type UserRepo struct {
	DB *DB
}

func (ur *UserRepo) CreateUser(c context.Context, user twitter.User) (twitter.User, error) {
	panic("implement me ")
}
func (ur *UserRepo) GetByUsername(ctx context.Context, username string) (twitter.User, error) {
	query := `SELECT * FROM users WHERE id = $1 LIMIT 1`
	u := twitter.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query); err != nil {
		if pgxscan.NotFound(err) {
			return twitter.User{}, twitter.ErrNotFound
		}
		return twitter.User{}, fmt.Errorf("error select: %v", err)
	}
	return u, nil
}
func (ur *UserRepo) GetByEmail(c context.Context, email string) (twitter.User, error) {
	query := `SELECT * FROM users WHERE email = $1 LIMIT 1`
	u := twitter.User{}

	if err := pgxscan.Get(c, ur.DB.Pool, &u, query); err != nil {
		if pgxscan.NotFound(err) {
			return twitter.User{}, twitter.ErrNotFound
		}
		return twitter.User{}, fmt.Errorf("error select: %v", err)
	}
	return u, nil
}
