//go:generate go run github.com/vektah/dataloaden UserLoader string twitter/graph.*User

package graph

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"twitter"
)

var loadersKey = "dataloaders"

type Loaders struct {
	UserByID UserLoader
}

type Repos struct {
	UserRepo twitter.UserRepo
}

func DataLoaderMiddleware(repos *Repos) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
				UserByID: UserLoader{
					wait:     1 * time.Millisecond,
					maxBatch: 100,
					fetch: func(ids []string) ([]User, []error) {
						users, err := repos.UserRepo.GetByIds(r.Context(), ids)
						if err != nil {
							return nil, []error{err}
						}

						userById := map[string]User{}

						for _, u := range users {
							userById[u.ID] = *mapUser(u)
						}

						result := make([]User, len(ids))

						for i, id := range ids {
							user, ok := userById[id]
							if !ok {
								return nil, []error{fmt.Errorf("user with id: %s is missing", id)}
							}
							result[i] = user
						}
						return result, nil
					},
				},
			})

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})
	}
}

func DataLoaderFor(c context.Context) *Loaders {

	return c.Value(loadersKey).(*Loaders)
}
