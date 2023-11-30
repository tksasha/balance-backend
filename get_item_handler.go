package main

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/tksasha/balance/model"
	"github.com/valyala/fasthttp"
)

func GetItemHandler(db *sql.DB) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		id, err := strconv.Atoi(ctx.UserValue("id").(string))
		if err != nil {
			JSON(StatusBadRequest, ctx, model.NewErrors("id", "is invalid")) // TODO: use constant here

			return
		}

		item, err := FindItem(db, id)

		if errors.Is(err, RecordNotFoundError) {
			JSON(StatusNotFound, ctx, model.NewErrors("base", "resource not found")) // TODO: use constant here

			return
		}

		if err != nil {
			JSON(StatusInternalServerError, ctx, model.NewErrors("base", "Internal Server Error")) // TODO: use constant here

			return
		}

		JSON(StatusOK, ctx, &item)
	}
}
