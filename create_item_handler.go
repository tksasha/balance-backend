package main

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

func parseParams(body []byte) (*ItemParams, error) {
	params := new(ItemParams)

	if err := json.Unmarshal(body, params); err != nil {
		return nil, err
	}

	return params, nil
}

func CreateItemHandler(ctx *fasthttp.RequestCtx) {
	db := Open()

	defer Close(db)

	params, err := parseParams(ctx.PostBody())
	if err != nil {
		UnprocessableEntity(ctx, err)

		return
	}

	item, err := CreateItem(db, params)
	if err == nil {
		Created(ctx, &item)

		return
	}

	if errors.Is(err, ClientError) {
		UnprocessableEntity(ctx, item.Errors)
	} else {
		InternalServerError(ctx, err)
	}
}
