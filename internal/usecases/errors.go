package usecases

import "errors"

var (
	ErrNotFound = errors.New("[APP ERROR] NOT FOUND")
	ErrUnknown  = errors.New("[APP ERROR] UNKNOWN")
)
