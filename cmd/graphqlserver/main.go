package main

import (
	"context"
	"fmt"
	"log"
	"twitter/config"
	"twitter/postgres"
)

func main() {

	ctx := context.Background()
	conf := config.New()
	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working ...")
}
