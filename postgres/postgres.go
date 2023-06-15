package postgres

import (
	"context"
	"log"
	"twitter/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(c context.Context, config *config.Config) *DB {
	dbConf, err := pgxpool.ParseConfig(config.Database.URL)
	if err != nil {
		log.Fatalf("can't parse postgres config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(c, dbConf)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	db := &DB{Pool: pool}
	db.Pool.Ping(c)

	return db
}

func (db *DB) Open(c context.Context) {
	if err := db.Pool.Ping(c); err != nil {
		log.Fatalf("can't ping postgres: %v", err)
	}
}

func (db *DB) Close() {
	db.Pool.Close()
}
