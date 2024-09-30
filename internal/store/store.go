package store

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Store struct {
	db       *sql.DB
	migrator Migrator
}

type Migrator interface {
	ApplyMigrations(db *sql.DB) error
}

func NewStore(databaseURI string, migrator Migrator) *Store {
	store := &Store{}

	db, err := sql.Open("pgx", databaseURI)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute)

	store.db = db
	store.migrator = migrator

	return store

}
