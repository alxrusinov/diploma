package customerrors

type NoOrderError struct {
	Err error
}

func (err *NoOrderError) Unwrap() error {
	return err.Err
}

func (err *NoOrderError) Error() string {
	return err.Err.Error()
}
