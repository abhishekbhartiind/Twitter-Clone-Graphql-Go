package domain

import (
	"context"
	"log"
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

func (ts *TweetService) Delete(c context.Context, id string) error {
	currentId, err := twitter.GetUserIdFromContext(c)
	if err != nil {
		return twitter.ErrUnAuthenticate
	}

	tweet, err := ts.TweetRepo.GetById(c, id)
	if err != nil {
		log.Println("same error is here :", err, ":::", currentId)
		return err
	}
	if !tweet.CanDelete(twitter.User{ID: currentId}) {
		return twitter.ErrForbidden
	}

	return ts.TweetRepo.Delete(c, id)
}

func (ts *TweetService) CreateReply(c context.Context, parentID string, input twitter.CreateTweetInput) (twitter.Tweet, error) {

	currentId, err := twitter.GetUserIdFromContext(c)
	if err != nil {
		return twitter.Tweet{}, twitter.ErrUnAuthenticate
	}

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return twitter.Tweet{}, twitter.ErrValidation
	}

	if _, err := ts.TweetRepo.GetById(c, parentID); err != nil {
		return twitter.Tweet{}, twitter.ErrNotFound
	}

	tweet, err := ts.TweetRepo.Create(c, twitter.Tweet{
		UserID:   currentId,
		Body:     input.Body,
		ParentId: &parentID,
	})

	if err != nil {
		return twitter.Tweet{}, err
	}

	return tweet, nil
}

func (ts *TweetService) GetAllReplyTweet(c context.Context, parentID string) ([]twitter.Tweet, error) {
	return ts.TweetRepo.GetAllReplyTweet(c, parentID)
}
