package router

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func New() func(*fasthttp.RequestCtx) {
	r := router.New()

	handler := func() func(ctx *fasthttp.RequestCtx) {
		log.Println("Server is started")

		return func(ctx *fasthttp.RequestCtx) {
			log.Println("you middlewares might be places here")

			r.Handler(ctx)
		}
	}

	return handler()
}
