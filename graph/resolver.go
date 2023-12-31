package graph

import (
	"context"
	"errors"
	"net/http"
	"twitter"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	AuthService  twitter.AuthService
	TweetService twitter.TweetService
	UserService  twitter.UserServices
}

type queryResolver struct {
	*Resolver
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	*Resolver
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type tweetResolver struct {
	*Resolver
}

func (r *Resolver) Tweet() TweetResolver {
	return &tweetResolver{r}
}

func buildBadRequest(c context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(c),
		Extensions: map[string]interface{}{
			"code": http.StatusBadRequest,
		},
	}
}

func buildUnAuthenticated(c context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(c),
		Extensions: map[string]interface{}{
			"code": http.StatusUnauthorized,
		},
	}
}

func buildForbidden(c context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(c),
		Extensions: map[string]interface{}{
			"code": http.StatusForbidden,
		},
	}
}

func buildNotFound(c context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(c),
		Extensions: map[string]interface{}{
			"code": http.StatusNotFound,
		},
	}
}

func buildError(c context.Context, err error) error {

	switch {
	case errors.Is(err, twitter.ErrNotFound):
		return buildNotFound(c, err)
	case errors.Is(err, twitter.ErrForbidden):
		return buildForbidden(c, err)
	case errors.Is(err, twitter.ErrUnAuthenticate):
		return buildUnAuthenticated(c, err)
	case errors.Is(err, twitter.ErrValidation):
		return buildBadRequest(c, err)

	default:
		return err
	}
}
