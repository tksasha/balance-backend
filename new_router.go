package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func NewRouter() func(ctx *fasthttp.RequestCtx) {
	r := router.New()

	r.POST("/items", CreateItemHandler)
	r.GET("/items/{id}", GetItemHandler)

	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType(MIMEApplicationJSON)

		r.Handler(ctx)
	}
}
