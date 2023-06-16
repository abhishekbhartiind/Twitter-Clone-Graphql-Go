package postgres

import (
	"context"
	"fmt"
	"log"
	"path"
	"runtime"
	"twitter/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
	conf *config.Config
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

	db := &DB{Pool: pool, conf: config}
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

func (db *DB) Migrate() error {

	_, f, _, _ := runtime.Caller(0)

	migrationPath := fmt.Sprintf("file:///%s/migrations", path.Dir(f))

	m, err := migrate.New(migrationPath, db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("error while updating the migrate instance %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error while migrate up %v", err)
	}
	return nil
}
