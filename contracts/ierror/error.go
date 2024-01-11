package ierror

import (
	"errors"

	"golang.org/x/xerrors"
)

type Error struct {
	xerrors.Wrapper

	message string
	fields  []any
}

func New(message string, fields ...any) *Error {
	return &Error{
		message: message,
		fields:  fields,
	}
}

func (e Error) Error() string {
	return e.message + "|" + fieldsBuilder(e.fields)
}

func (e Error) Unwrap() error {
	return errors.New(e.message)
}
