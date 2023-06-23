package graph

import "twitter"

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
