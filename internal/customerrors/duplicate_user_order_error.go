package customerrors

type DuplicateUserOrderError struct {
	Err error
}

func (err *DuplicateUserOrderError) Unwrap() error {
	return err.Err
}

func (err *DuplicateUserOrderError) Error() string {
	return err.Err.Error()
}
