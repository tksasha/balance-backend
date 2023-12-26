package api

import (
	"fmt"
)

type Error struct {
	err error
}

func NewError(err error) *Error {
	return &Error{err}
}

func (err *Error) Error() string {
	return fmt.Sprintf(`{"api":%v}`, err.err)
}

func (err *Error) MarshalJSON() ([]byte, error) {
	return []byte(err.Error()), nil
}
