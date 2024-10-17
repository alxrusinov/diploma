package customerrors

type CustomError interface {
	Unwrap() error
	Error() string
}
