package store

import (
	"database/sql"
	"log"
)

type DBStore struct {
	db *sql.DB
}

func CreateDBStore(databaseURI string) Store {
	store := &DBStore{}

	db, err := sql.Open("pgx", databaseURI)

	if err != nil {
		log.Fatal(err)
	}

	store.db = db

	return store

}
