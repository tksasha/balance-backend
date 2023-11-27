package main

import (
	"errors"
)

var (
	ClientError = errors.New("ClientError")
	ServerError = errors.New("ServerError")
)

const (
	StatusOK                  = 200 // 200
	StatusCreated             = 201
	StatusNotFound            = 404 // 400
	StatusUnprocessableEntity = 422
	StatusInternalServerError = 500 // 500

	MIMEApplicationJSON = "application/json"
)
