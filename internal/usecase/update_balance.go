package usecase

func (usecase *Usecase) UpdateBalance(balance int, userID string) error {
	err := usecase.store.UpdateBalance(balance, userID)

	return err
}
