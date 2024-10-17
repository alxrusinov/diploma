package store

import (
	"errors"
	"time"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) SetWithdrawls(withdrawn *model.Withdrawn, userID string) error {
	tx, err := store.db.Begin()

	if err != nil {
		return err
	}

	withdrawn.Round()

	balance := new(model.Balance)

	err = tx.QueryRow(selectBalanceQuery, userID).Scan(&balance.Current)

	if err != nil {
		tx.Rollback()
		return err
	}

	balance.Round()

	if !withdrawn.IsWithdrawAvailable(balance.Current) {
		tx.Rollback()
		return &customerrors.PaymentRequiredError{Err: errors.New("withdraw is unavailable")}
	}

	_, err = tx.Exec(updateBalanceQuery, balance.Current-withdrawn.Sum, userID)

	if err != nil {
		tx.Rollback()
		return err
	}

	withdrawn.ProcessedAt = time.Now().Format(time.RFC1123Z)

	_, err = tx.Exec(setWithdrawnQuery, userID, withdrawn.Order, withdrawn.Sum, withdrawn.ProcessedAt)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
