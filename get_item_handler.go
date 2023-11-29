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
			errs := model.NewErrors()

			errs.Add("id", "is invalid")

			JSON(StatusBadRequest, ctx, errs)

			return
		}

		item, err := GetItemQuery(db, id)

		if errors.Is(err, RecordNotFoundError) {
			errs := model.NewErrors()

			errs.Add("base", "resource not found")

			JSON(StatusNotFound, ctx, errs)

			return
		}

		if err != nil {
			errs := model.NewErrors()

			errs.Add("base", "Internal Server Error")

			JSON(StatusInternalServerError, ctx, errs)

			return
		}

		JSON(StatusOK, ctx, &item)
	}
}
