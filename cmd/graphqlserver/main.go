package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"twitter/config"
	"twitter/domain"
	"twitter/graph"
	"twitter/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	ctx := context.Background()
	conf := config.New()
	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	// SERVICES
	userRepo := postgres.NewUserRepo(db)

	authService := domain.NewAuthService(userRepo)

	router.Handle("/", playground.Handler("twitter clone", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			AuthService: authService,
		}}),
	))

	log.Fatal(http.ListenAndServe(":8080", router))

}
