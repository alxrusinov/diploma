package store

import (
	"database/sql"
	"log"

	"github.com/alxrusinov/diploma/internal/migrator"
	"github.com/alxrusinov/diploma/internal/model"
)

type DBStore struct {
	db       *sql.DB
	migrator *migrator.Migrator
}

func (store *DBStore) FindUserByLogin(user *model.User) (bool, error) {
	return true, nil
}

func (store *DBStore) FindUserByLoginPassword(user *model.User) (bool, error) {
	return true, nil
}

func (store *DBStore) CreateUser(user *model.User) (bool, error) {
	return true, nil
}

func (store *DBStore) UpdateUser(token *model.Token) (*model.Token, error) {
	return new(model.Token), nil
}

func (store *DBStore) AddOrder(order *model.Order) (bool, error) {
	return true, nil
}

func (store *DBStore) GetOrders(login string) ([]model.OrderResponse, error) {
	return make([]model.OrderResponse, 0), nil
}

func (store *DBStore) RunMigration() error {
	err := store.migrator.ApplyMigrations(store.db)

	return err
}

func CreateDBStore(databaseURI string, migrator *migrator.Migrator) Store {
	store := &DBStore{}

	db, err := sql.Open("pgx", databaseURI)

	if err != nil {
		log.Fatal(err)
	}

	store.db = db
	store.migrator = migrator

	return store

}
