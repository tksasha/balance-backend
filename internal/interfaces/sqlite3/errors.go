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

type QueryError struct {
	err error
}

func NewQueryError(err error) QueryError {
	return QueryError{err}
}

func (err QueryError) Error() string {
	return fmt.Sprintf(`{"database":{"query":"%s"}}`, err.err)
}

func (err QueryError) MarshalJSON() ([]byte, error) {
	return []byte(err.Error()), nil
}
