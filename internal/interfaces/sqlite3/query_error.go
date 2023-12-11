package sqlite3

import (
	"fmt"
	"strings"
)

type QueryError struct {
	title       string
	description string
}

func NewQueryError(description error) *QueryError {
	return &QueryError{
		title:       "[DB][QUERY]",
		description: strings.ToUpper(description.Error()),
	}
}

func (err *QueryError) Error() string {
	return fmt.Sprintf("%s %s", err.title, err.description)
}
