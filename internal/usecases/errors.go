package usecases

import (
	"fmt"
)

type UnknownError struct {
	err error
}

func NewUnknownError(err error) *UnknownError {
	return &UnknownError{err}
}

func (err *UnknownError) Error() string {
	return fmt.Sprintf(`{"application":{"UnknownError":"%s"}}`, err.err)
}

func (err *UnknownError) MarshalJSON() ([]byte, error) {
	return []byte(err.Error()), nil
}

type NotFoundError struct {
	err error
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{err}
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf(`{"application":{"NotFoundError":"%s"}}`, err.err)
}

func (err *NotFoundError) MarshalJSON() ([]byte, error) {
	return []byte(err.Error()), nil
}
