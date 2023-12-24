package usecases

import (
	"fmt"
)

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
