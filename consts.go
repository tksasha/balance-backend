package main

import (
	"errors"
)

var (
	ClientError = errors.New("ClientError")
	ServerError = errors.New("ServerError")

	RecordNotFoundError = errors.New("RecordNotFoundError")
)

const (
	StatusOK                  = 200 // 200
	StatusCreated             = 201
	StatusBadRequest          = 400 // 400
	StatusNotFound            = 404
	StatusUnprocessableEntity = 422
	StatusInternalServerError = 500 // 500

	MIMEApplicationJSON = "application/vnd.api+json"

	ErrInvalid  = "invalid"
	ErrNotFound = "not found"
	ErrServer   = "server error"
)
