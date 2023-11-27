package main

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

func CreateItemHandler(ctx *fasthttp.RequestCtx) {
	db := Open()

	defer Close(db)

	params := new(ItemParams)

	if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
		ctx.SetStatusCode(StatusUnprocessableEntity)

		json.NewEncoder(ctx).Encode(err)

		return
	}

	item, err := CreateItem(db, params)
	if err == nil {
		ctx.SetStatusCode(StatusCreated)

		json.NewEncoder(ctx).Encode(&item)
	} else {
		if errors.Is(err, ClientError) {
			ctx.SetStatusCode(StatusUnprocessableEntity)

			json.NewEncoder(ctx).Encode(item.Errors)
		} else {
			ctx.SetStatusCode(StatusInternalServerError)

			json.NewEncoder(ctx).Encode(err)
		}
	}
}
