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

type ErrorResp struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func RespError(ctx *fasthttp.RequestCtx, err error) {
	RespErrorWithCodeAndMessage(ctx, nil, err.Error(), fasthttp.StatusInternalServerError)
}

func RespErrorWithCodeAndMessage(ctx *fasthttp.RequestCtx, errorCode *string, errorMessage string, statusCode int) {

	res := &ErrorResp{
		Message: errorMessage,
	}
	if errorCode != nil {
		res.Code = *errorCode
	}

	if err := json.NewEncoder(ctx).Encode(res); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	} else {
		ctx.SetStatusCode(statusCode)
	}
}
