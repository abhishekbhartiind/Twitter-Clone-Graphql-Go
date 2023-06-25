package twitter

import (
	"context"
)

type contextKey string

var (
	contextAuthIdKey contextKey = "currentUserId"
)

func GetUserIdFromContext(c context.Context) (string, error) {

	if c.Value(contextAuthIdKey) == nil {
		return "", ErrNoUserInContext
	}

	userId, ok := c.Value(contextAuthIdKey).(string)
	if !ok {
		return "", ErrNoUserInContext
	}

	return userId, nil
}

func PutUserIdIntoContext(c context.Context, id string) context.Context {
	return context.WithValue(c, contextAuthIdKey, id)
}
