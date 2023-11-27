package main

import (
	"errors"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
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

func main() {
	r := router.New()

	r.POST("/items", CreateItemHandler)
	r.GET("/items/{id}", GetItemHandler)

	log.Fatal(fasthttp.ListenAndServe(":3000", r.Handler))
}
