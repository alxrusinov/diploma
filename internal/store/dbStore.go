package store

import (
	"database/sql"
	"log"

	"github.com/alxrusinov/diploma/internal/model"
)

type DBStore struct {
	db *sql.DB
}

func (store *DBStore) CheckUserExists(user *model.User) (bool, error) {
	return false, nil
}

func (store *DBStore) CreateUser(user *model.User) error {
	return nil
}

func (store *DBStore) UpdateUser(toke *model.Token) (*model.Token, error) {
	return &model.Token{}, nil
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
