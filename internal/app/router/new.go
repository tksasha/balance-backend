package router

import (
	"database/sql"
	"log"

	"github.com/fasthttp/router"
	"github.com/tksasha/balance/internal/controller"
	"github.com/valyala/fasthttp"
)

func New(db *sql.DB) func(*fasthttp.RequestCtx) {
	r := router.New()

	handler := func() func(ctx *fasthttp.RequestCtx) {
		log.Println("Server is started")

		return func(ctx *fasthttp.RequestCtx) {
			log.Println("you middlewares might be places here")

			r.Handler(ctx)
		}
	}

	items := controller.NewItemController(db)
	r.GET("/items/{id}", items.Show)

	categories := controller.NewCategoryController(db)
	r.POST("/categories", categories.Create)

	return handler()
}
