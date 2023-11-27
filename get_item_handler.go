package main

import (
	"database/sql"
	"errors"
	"strconv"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func GetItemHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType(MIMEApplicationJSON) // TODO: move to middleware

	db := Open()

	defer Close(db)

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(StatusNotFound)

		json.NewEncoder(ctx).Encode("")

		return
	}

	item, err := GetItemQuery(db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.SetStatusCode(StatusNotFound)

			json.NewEncoder(ctx).Encode("")

			return
		} else {
			ctx.SetStatusCode(StatusInternalServerError)

			json.NewEncoder(ctx).Encode(err)

			return
		}
	}

	ctx.SetStatusCode(StatusOK)

	json.NewEncoder(ctx).Encode(item)
}
