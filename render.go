package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func JSON(statusCode int, ctx *fasthttp.RequestCtx, body any) {
	ctx.SetStatusCode(statusCode)

	ctx.SetContentType(MIMEApplicationJSON)

	json.NewEncoder(ctx).Encode(body)
}
