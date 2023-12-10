package usecases

import "errors"

var (
	ErrNotFound = errors.New("NOT FOUND")
	ErrUnknown  = errors.New("UNKNOWN ERROR")
)
