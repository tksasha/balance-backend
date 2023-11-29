package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/tksasha/balance/jsonapi"
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

func GetItemHandler(db *sql.DB) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		id, err := strconv.Atoi(ctx.UserValue("id").(string))
		if err != nil {
			JSON(StatusBadRequest, ctx, jsonapi.NewErrors("parameter", "id", ErrInvalid))

			return
		}

		item, err := GetItemQuery(db, id)

		if errors.Is(err, RecordNotFoundError) {
			JSON(StatusNotFound, ctx, jsonapi.NewErrors("parameter", "id", ErrNotFound))

			return
		}

		if err != nil {
			JSON(StatusInternalServerError, ctx, jsonapi.NewErrors(ErrServer))

			return
		}

		JSON(StatusOK, ctx, &item)
	}
}
