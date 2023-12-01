package main

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

func CreateItemHandler(db *sql.DB) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		params := new(ItemParams)

		if err := json.Unmarshal(ctx.PostBody(), params); err != nil {
			JSON(fasthttp.StatusUnprocessableEntity, ctx, err)

			return
		}

		item, err := CreateItem(db, params)
		if err == nil {
			JSON(fasthttp.StatusCreated, ctx, &item)

			return
		}

		switch {
		case errors.Is(err, RecordInvalidError):
			JSON(fasthttp.StatusUnprocessableEntity, ctx, item.Errors)
		case errors.Is(err, InternalServerError):
			JSON(fasthttp.StatusInternalServerError, ctx, err)
		default:
			panic("UNKNOWN ERROR !!!")
		}
	}
}
