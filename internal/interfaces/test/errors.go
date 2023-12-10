package test

import "errors"

var (
	ErrNotFound = errors.New("NOT FOUND")
	ErrDB       = errors.New("DB ERROR")
)
