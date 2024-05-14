package router

import (
	"github.com/valyala/fasthttp"
	"gitlab.com/miguelit0/toDoApp/internal/helpers"
)

func (r *RouteHandlers) routesToDoApp() []Route {
	return []Route{
		{
			Name:        "Create task",
			Method:      fasthttp.MethodPost,
			Pattern:     "/tasks",
			HandlerFunc: r.CreateTask(),
		},
	}
}

func (r *RouteHandlers) CreateTask() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		message := "Task created"
		helpers.RespOk(ctx, message)

	}
}
