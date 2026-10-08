package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	Message string
	Err     error
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

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return e.Message == t.Message && (t.Err == nil || errors.Is(e.Err, t.Err))
}
