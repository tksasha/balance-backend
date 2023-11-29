package main

import (
	"errors"
)

var (
	ClientError = errors.New("ClientError") // TODO: DELME
	ServerError = errors.New("ServerError") // TODO: DELME

	RecordNotFoundError = errors.New("RecordNotFoundError")
	RecordInvalidError  = errors.New("RecordInvalidError")
	InternalServerError = errors.New("InternalServerError")
)

const (
	StatusOK                  = 200 // 200
	StatusCreated             = 201
	StatusBadRequest          = 400 // 400
	StatusNotFound            = 404
	StatusUnprocessableEntity = 422
	StatusInternalServerError = 500 // 500

	MIMEApplicationJSON = "application/json"
)
