package customerrors

type PaymentRequiredError struct {
	Err error
}

func (err *PaymentRequiredError) Unwrap() error {
	return err.Err
}

func (err *PaymentRequiredError) Error() string {
	return err.Err.Error()
}
