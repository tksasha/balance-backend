package sqlite3

import (
	"fmt"
)

type GetIDError struct {
	err error
}

func NewGetIDError(err error) error {
	return GetIDError{err}
}

func (err GetIDError) Error() string {
	return fmt.Sprintf(`{"database":{"GetIDError":"%s"}}`, err.err)
}

func (err GetIDError) MarshalJSON() ([]byte, error) {
	return []byte(err.Error()), nil
}
