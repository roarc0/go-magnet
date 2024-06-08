package magnet

// ErrorInvalidMagnet represents an error when a magnet URI is invalid.
type ErrorInvalidMagnet struct {
	msg string
	err error
}

func (e *ErrorInvalidMagnet) Unwrap() error {
	return e.err
}

func (e *ErrorInvalidMagnet) Error() string {
	return e.msg
}
