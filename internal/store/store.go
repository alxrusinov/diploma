package store

import (
	"database/sql"
	"log"

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

	store.db = db
	store.migrator = migrator

	return store

}
