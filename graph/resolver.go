package graph

import (
	"context"
	"net/http"
	"twitter"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	AuthService twitter.AuthService
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

func buildBadRequest(c context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(c),
		Extensions: map[string]interface{}{
			"code": http.StatusBadRequest,
		},
	}
}
