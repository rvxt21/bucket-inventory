package errors

import "fmt"

type Error struct {
	Message string
	Err     error
}

func (e *Error) ExpandError(err error) {
	e.Err = err
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}

	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Wrap(err error) *Error {
	return &Error{
		Message: e.Message,
		Err:     err,
	}
}
