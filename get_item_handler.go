package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"
)

func GetItemHandler(ctx *fasthttp.RequestCtx) {
	db := Open()

	defer Close(db)

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		NotFound(ctx, "")

		return
	}

	item, err := GetItemQuery(db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFound(ctx, "")

			return
		} else {
			ctx.SetStatusCode(StatusInternalServerError)

			json.NewEncoder(ctx).Encode(err)

			return
		}
	}

	OK(ctx, &item)
}
