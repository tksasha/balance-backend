package main

import (
	"errors"
)

var (
	RecordNotFoundError = errors.New("RecordNotFoundError")
	RecordInvalidError  = errors.New("RecordInvalidError")
	InternalServerError = errors.New("InternalServerError")
)
