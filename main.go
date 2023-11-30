package main

import (
	"log"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	db := Open()
	defer Close(db)

	time.Local = time.UTC // TODO: do something

	r := router.New()

	// r.POST("/items", CreateItemHandler)
	r.GET("/items/{id}", GetItemHandler(db))

	handler := func() func(ctx *fasthttp.RequestCtx) {
		log.Println("Server is started")

		return func(ctx *fasthttp.RequestCtx) {
			log.Println("your middlewares might be placed here")

			r.Handler(ctx)
		}
	}

	log.Fatal(fasthttp.ListenAndServe(":3000", handler()))
}
