package sqlite3

import (
	"fmt"
	"strings"
)

type IDError struct {
	title       string
	description string
}

func NewIDError(description error) *IDError {
	return &IDError{
		title:       "[DB][ID]",
		description: strings.ToUpper(description.Error()),
	}
}

func (err *IDError) Error() string {
	return fmt.Sprintf("%s %s", err.title, err.description)
}
