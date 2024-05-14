package helpers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func RespOk(ctx *fasthttp.RequestCtx, res interface{}) {
	ctx.SetContentType("application/json; charset=UTF-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(res); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}
