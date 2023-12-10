package test

import "errors"

var (
	ErrNotFound = errors.New("[DB ERROR] NOT FOUND")
	ErrUnknown  = errors.New("[DB ERROR] UNKNOWN")
)
