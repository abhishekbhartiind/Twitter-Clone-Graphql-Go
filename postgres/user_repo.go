package postgres

import (
	"context"
	"fmt"
	"twitter"

	"github.com/georgysavva/scany/pgxscan"
	// "github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v4"
)

type UserRepo struct {
	DB *DB
}

func (ur *UserRepo) CreateUser(c context.Context, user twitter.User) (twitter.User, error) {

	tx, err := ur.DB.Pool.Begin(c)
	if err != nil {
		return twitter.User{}, fmt.Errorf("err while begin transaction %v", err)
	}
	defer tx.Rollback(c)

	user, err = createUser(c, tx, user)
	if err != nil {
		return twitter.User{}, err
	}

	if err := tx.Commit(c); err != nil {
		return twitter.User{}, fmt.Errorf("error while commiting %v", err)
	}

	return user, nil
}

func createUser(c context.Context, tx pgx.Tx, user twitter.User) (twitter.User, error) {
	query := `INSERT INTO users (username, email,password) VALUES ($1, $2, $3) RETURNING *;`

	u := twitter.User{}

	if err := pgxscan.Get(c, tx, &u, query, user.Username, user.Email, user.Password); err != nil {
		return twitter.User{}, fmt.Errorf("error while inserting %v", err)
	}
	return u, nil
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
