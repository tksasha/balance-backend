package sqlite3

import (
	"fmt"
)

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
