package usecase

func (usecase *Usecase) UpdateBalance(balance float32, userID string) error {
	err := usecase.store.UpdateBalance(balance, userID)

	return err
}
