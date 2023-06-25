package main

import (
	"net/http"
	"twitter"
)

func authMiddleware(authTokenService twitter.AuthTokenService) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			token, err := authTokenService.ParseTokenFromRequest(ctx, r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// adding user id into context
			ctx = twitter.PutUserIdIntoContext(ctx, token.Sub)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
