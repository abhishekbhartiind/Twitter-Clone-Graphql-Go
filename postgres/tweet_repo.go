package postgres

import (
	"context"
	"fmt"
	"log"
	"twitter"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type TweetRepo struct {
	DB *DB
}

func NewTweetRepo(db *DB) *TweetRepo {
	return &TweetRepo{
		DB: db,
	}
}

func (tr *TweetRepo) All(c context.Context) ([]twitter.Tweet, error) {
	return getAllTweets(c, tr.DB.Pool)
}

func getAllTweets(c context.Context, q pgxscan.Querier) ([]twitter.Tweet, error) {
	query := `SELECT * from tweets ORDER BY created_at DESC`

	t := []twitter.Tweet{}

	if err := pgxscan.Select(c, q, &t, query); err != nil {
		return nil, fmt.Errorf("error while getting all tweets: %v", err)
	}

	return t, nil
}

func (tr *TweetRepo) Create(c context.Context, tweet twitter.Tweet) (twitter.Tweet, error) {
	tx, err := tr.DB.Pool.Begin(c)
	if err != nil {
		return twitter.Tweet{}, fmt.Errorf("error while transaction %v", err)
	}
	defer tx.Rollback(c)
	tweet, err = createTweet(c, tx, tweet)
	if err != nil {
		return twitter.Tweet{}, err
	}

	if err := tx.Commit(c); err != nil {
		return twitter.Tweet{}, fmt.Errorf("error while commiting: %v", err)
	}

	return tweet, nil
}

func createTweet(c context.Context, tx pgx.Tx, tweet twitter.Tweet) (twitter.Tweet, error) {

	query := `INSERT INTO tweets (body, user_id, parent_id) VALUES ($1, $2, $3) RETURNING *;`

	t := twitter.Tweet{}

	if err := pgxscan.Get(c, tx, &t, query, tweet.Body, tweet.UserID, tweet.ParentId); err != nil {
		return twitter.Tweet{}, fmt.Errorf("error while tweet insert %v", err)
	}
	return t, nil
}

func (tr *TweetRepo) GetById(c context.Context, id string) (twitter.Tweet, error) {
	log.Println("testomg ", id)
	return getTweetById(c, tr.DB.Pool, id)
}

func getTweetById(c context.Context, q pgxscan.Querier, id string) (twitter.Tweet, error) {

	query := `SELECT * FROM tweets WHERE id = $1 LIMIT 1;`

	t := twitter.Tweet{}

	if err := pgxscan.Get(c, q, &t, query, id); err != nil {
		if pgxscan.NotFound(err) {
			return twitter.Tweet{}, twitter.ErrNotFound
		}

		return twitter.Tweet{}, fmt.Errorf("error get tweet: %+v", err)
	}

	return t, nil
}

func (tr *TweetRepo) Delete(c context.Context, id string) error {
	tx, err := tr.DB.Pool.Begin(c)
	if err != nil {
		return fmt.Errorf("error while transaction %v", err)
	}
	defer tx.Rollback(c)

	if err := deletTweet(c, tx, id); err != nil {
		return err
	}

	if err := tx.Commit(c); err != nil {
		return fmt.Errorf("error while commiting: %v", err)
	}

	return nil
}

func deletTweet(c context.Context, tx pgx.Tx, id string) error {

	query := `DELETE FROM tweets Where id = $1;`

	if _, err := tx.Exec(c, query, id); err != nil {
		return fmt.Errorf("error while tweet delete %v", err)
	}
	return nil
}
