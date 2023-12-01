package main

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"
)

func GetItemHandler(db *sql.DB) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		id, err := strconv.Atoi(ctx.UserValue("id").(string))
		if err != nil {
			JSON(fasthttp.StatusBadRequest, ctx, err)

			return
		}

		item, err := FindItem(db, id)

		if errors.Is(err, RecordNotFoundError) {
			JSON(fasthttp.StatusNotFound, ctx, err)

			return
		}

		if err != nil {
			JSON(fasthttp.StatusInternalServerError, ctx, err)

			return
		}

		JSON(fasthttp.StatusOK, ctx, &item)
	}
}
