package errors

import (
	"errors"
)

var (
	RecordInvalidError = errors.New("Record Invalid Error")
	UnknownError       = errors.New("Unknown Error")
)

const (
	CantBeBlank = "can't be blank"
)
