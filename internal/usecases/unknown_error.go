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
