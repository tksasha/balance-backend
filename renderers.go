package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func JSON(ctx *fasthttp.RequestCtx, body any) {
	ctx.SetContentType(MIMEApplicationJSON)

	json.NewEncoder(ctx).Encode(body)
}

// 200
func OK(ctx *fasthttp.RequestCtx, body any) {
	ctx.SetStatusCode(StatusOK)

	JSON(ctx, body)
}

// 201
func Created(ctx *fasthttp.RequestCtx, body any) {
	ctx.SetStatusCode(StatusCreated)

	JSON(ctx, body)
}

// 404
func NotFound(ctx *fasthttp.RequestCtx, body any) {
	ctx.SetStatusCode(StatusNotFound)

	JSON(ctx, body)
}

// 422
func UnprocessableEntity(ctx *fasthttp.RequestCtx, body any) {
	ctx.SetStatusCode(StatusUnprocessableEntity)

	JSON(ctx, body)
}

// 500
func InternalServerError(ctx *fasthttp.RequestCtx, body any) {
	ctx.SetStatusCode(StatusInternalServerError)

	JSON(ctx, body)
}
