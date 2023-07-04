package graph

import (
	"context"
	"errors"
	"twitter"
)

func mapTweet(t twitter.Tweet) *Tweet {
	return &Tweet{
		ID:        t.ID,
		Body:      t.Body,
		UserID:    t.UserID,
		CreatedAt: t.CreatedAt,
	}
}

func mapTweets(tweets []twitter.Tweet) []*Tweet {
	tt := make([]*Tweet, len(tweets))
	for i, t := range tweets {
		tt[i] = mapTweet(t)
	}
	return tt
}

func (q *queryResolver) Tweets(ctx context.Context) ([]*Tweet, error) {
	tweets, err := q.TweetService.All(ctx)
	if err != nil {
		return nil, err
	}
	return mapTweets(tweets), nil
}

func (m *mutationResolver) CreateTweet(ctx context.Context, input CreateTweetInput) (*Tweet, error) {
	tweet, err := m.TweetService.Create(ctx, twitter.CreateTweetInput{
		Body: input.Body,
	})
	if err != nil {
		if errors.Is(err, twitter.ErrUnAuthenticate) {
			return nil, buildUnAuthenticated(ctx, err)
		}
		return nil, err
	}
	return mapTweet(tweet), nil
}

func (m *mutationResolver) DeleteTweet(ctx context.Context, id string) (bool, error) {

	if err := m.TweetService.Delete(ctx, id); err != nil {
		return false, buildError(ctx, err)
	}

	return true, nil
}

func mapToUser(u User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (t *tweetResolver) User(c context.Context, obj *Tweet) (*User, error) {

	u, e := DataLoaderFor(c).UserByID.Load(obj.UserID)
	uu := mapToUser(u)
	return uu, e

	// user, err := t.UserService.GetById(c, obj.UserID)
	// if err != nil {
	// 	return nil, buildError(c, err)
	// }
	// return mapUser(user), nil
}
