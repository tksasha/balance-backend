package usecases

import (
	"fmt"
)

type ApplicationError struct {
	err error
}

func NewApplicationError(err error) *ApplicationError {
	return &ApplicationError{err}
}

func (err *ApplicationError) Error() string {
	return fmt.Sprintf(`{"application":%v}`, err.err)
}

func (err *ApplicationError) MarshalJSON() ([]byte, error) {
	return []byte(err.Error()), nil
}
