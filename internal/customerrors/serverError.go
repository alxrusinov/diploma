package customerrors

type ServerError struct {
	Err error
}

func (err *ServerError) Unwrap() error {
	return err.Err
}

func (err *ServerError) Error() string {
	return err.Err.Error()
}
