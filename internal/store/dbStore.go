package store

import (
	"database/sql"
	"errors"
	"io"
	"log"
	"time"

	"github.com/alxrusinov/diploma/internal/migrator"
	"github.com/alxrusinov/diploma/internal/model"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBStore struct {
	db       *sql.DB
	migrator *migrator.Migrator
}

func (store *DBStore) FindUserByLogin(user *model.User) (bool, error) {

	row := store.db.QueryRow(`SELECT id FROM users WHERE login = $1`, user.Login)

	var login string

	err := row.Scan(&login)

	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}

	if login == "" {
		return false, nil
	}

	return true, nil
}

func (store *DBStore) FindUserByLoginPassword(user *model.User) (bool, error) {
	row := store.db.QueryRow(`SELECT id FROM users WHERE login = $1 and password = $2`, user.Login, user.Password)

	var login string

	err := row.Scan(&login)

	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}

	if login == "" {
		return false, nil
	}

	return true, nil
}

func (store *DBStore) CreateUser(user *model.User) (bool, error) {
	tx, err := store.db.Begin()

	if err != nil {
		return false, err
	}
	queryUser := `INSERT INTO users (login, password, token)
				VALUES ($1, $2, $3)
				RETURNING id;`

	queryBalance := `INSERT INTO users (user_id, balance)
				VALUES ($1, $2);`

	var userID string

	err = tx.QueryRow(queryUser, user.Login, user.Password, "").Scan(&userID)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = tx.Exec(queryBalance, userID, 0)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}

func (store *DBStore) UpdateUser(token *model.Token) (*model.Token, error) {
	query := `UPDATE users SET token = $1 WHERE login = $2`

	tokenRows, err := store.db.Query(query, token.Token, token.UserName)

	if err != nil || tokenRows.Err() != nil {
		return nil, err
	}

	return token, nil
}

func (store *DBStore) AddOrder(order *model.Order, login string) (bool, error) {
	tx, err := store.db.Begin()

	if err != nil {
		return false, err
	}

	var userID string

	err = tx.QueryRow(`SELECT id FROM users WHERE login = $1`, login).Scan(&userID)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	row, err := tx.Query(`INSERT INTO orders (user_id, number, process, accrual, uploaded_at)
	VALUES ($1, $2, $3, $4, $5)`, userID, order.Number, order.Process, order.Accrual, time.Now().Format(time.RFC3339))

	if err != nil || row.Err() != nil {
		tx.Rollback()
		return false, err
	}

	var sum int

	err = tx.QueryRow(`SELECT balance FROM balance WHERE user_id = $1`, userID).Scan(&sum)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	rows, err := tx.Query(`UPDATE balance SET balance = $1 WHERE user_id = $2`, sum+order.Accrual, userID)

	if err != nil || rows.Err() != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}

func (store *DBStore) GetOrders(login string) ([]model.OrderResponse, error) {
	res := make([]model.OrderResponse, 0)
	var userID string

	err := store.db.QueryRow(`SELECT id FROM users WHERE login = $1`, login).Scan(&userID)

	if err != nil {
		return res, err
	}

	rows, err := store.db.Query(`SELECT number, process, accrual, uploaded_at FROM orders WHERE user_id = $1`, userID)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		item := new(model.OrderResponse)
		err = rows.Scan(&item.Number, &item.Status, &item.Accrual, &item.UploadedAt)

		if err != nil {
			return res, err
		}

		res = append(res, *item)

	}

	if rows.Err() != nil {
		return res, err
	}

	return res, nil
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
