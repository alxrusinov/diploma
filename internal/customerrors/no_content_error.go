package customerrors

type NoContentError struct {
	Err error
}

func (err *NoContentError) Unwrap() error {
	return err.Err
}

func (err *NoContentError) Error() string {
	return err.Err.Error()
}
