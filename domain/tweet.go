package domain

import (
	"context"
	"twitter"
)

type TweetService struct {
	TweetRepo twitter.TweetRepo
}

func NewTweetService(tr twitter.TweetRepo) *TweetService {
	return &TweetService{
		TweetRepo: tr,
	}
}
func (ts *TweetService) All(c context.Context) ([]twitter.Tweet, error) {
	return ts.TweetRepo.All(c)
}

func (ts *TweetService) Create(c context.Context, input twitter.CreateTweetInput) (twitter.Tweet, error) {
	currentUserID, err := twitter.GetUserIdFromContext(c)
	if err != nil {
		return twitter.Tweet{}, twitter.ErrUnAuthenticate
	}

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return twitter.Tweet{}, err
	}
	tweet, err := ts.TweetRepo.Create(c, twitter.Tweet{
		Body:   input.Body,
		UserID: currentUserID,
	})
	if err != nil {
		return twitter.Tweet{}, err
	}
	return tweet, nil
}

func (ts *TweetService) GetById(c context.Context, id string) (twitter.Tweet, error) {
	return ts.TweetRepo.GetById(c, id)
}
