package customerrors

type DuplicateOwnerOrderError struct {
	Err error
}

func (err *DuplicateOwnerOrderError) Unwrap() error {
	return err.Err
}

func (err *DuplicateOwnerOrderError) Error() string {
	return err.Err.Error()
}
