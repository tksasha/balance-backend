package main

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

func parseParams(body []byte) (*ItemParams, error) {
	params := new(ItemParams)

	if err := json.Unmarshal(body, params); err != nil {
		return nil, err // TODO: prettify this error
	}

	return params, nil
}

func CreateItemHandler(ctx *fasthttp.RequestCtx) {
	db := Open()

	defer Close(db)

	params, err := parseParams(ctx.PostBody())
	if err != nil {
		JSON(StatusUnprocessableEntity, ctx, err) // TODO: prettify this error

		return
	}

	item, err := CreateItem(db, params)
	if err == nil {
		JSON(StatusCreated, ctx, &item)

		return
	}

	if errors.Is(err, ClientError) {
		JSON(StatusUnprocessableEntity, ctx, item.Errors)
	} else {
		JSON(StatusInternalServerError, ctx, err) // TODO: prettify this error
	}
}
