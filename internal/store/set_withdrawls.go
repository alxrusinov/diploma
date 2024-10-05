package store

import "github.com/alxrusinov/diploma/internal/model"

func (store *Store) SetWithdrawls(withdrawn *model.Withdrawn, userID string) error {
	withdrawn.Round()
	_, err := store.db.Exec(setWithdrawnQuery, userID, withdrawn.Order, withdrawn.Sum, withdrawn.ProcessedAt)

	return err
}
