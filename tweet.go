package twitter

import (
	"context"
	"fmt"
	"strings"
	"time"
)

var (
	tweetMinLength = 2
	tweetMaxLength = 250
)

func (t Tweet) CanDelete(user User) bool {
	return t.UserID == user.ID
}

type CreateTweetInput struct {
	Body string
}

func (in *CreateTweetInput) Sanitize() {
	in.Body = strings.TrimSpace(in.Body)
}

func (in *CreateTweetInput) Validate() error {

	if len(in.Body) < tweetMinLength {
		return fmt.Errorf("%w: body not enough, (%d) characters atleast", ErrValidation, tweetMinLength)
	}

	if len(in.Body) > tweetMaxLength {
		return fmt.Errorf("%w: body too long, (%d) characters at max, ", ErrValidation, tweetMaxLength)
	}

	return nil
}

type Tweet struct {
	ID        string
	Body      string
	UserID    string
	ParentId  *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TweetService interface {
	All(c context.Context) ([]Tweet, error)
	Create(c context.Context, input CreateTweetInput) (Tweet, error)
	CreateReply(c context.Context, parentID string, input CreateTweetInput) (Tweet, error)
	GetAllReplyTweet(c context.Context, parentID string) ([]Tweet, error)
	GetById(c context.Context, id string) (Tweet, error)
	Delete(c context.Context, id string) error
}

type TweetRepo interface {
	All(c context.Context) ([]Tweet, error)
	Create(c context.Context, tweet Tweet) (Tweet, error)
	GetById(c context.Context, id string) (Tweet, error)
	GetAllReplyTweet(c context.Context, parentID string) ([]Tweet, error)
	Delete(c context.Context, id string) error
}
