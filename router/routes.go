package router

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"gitlab.com/miguelit0/toDoApp/internal/helpers"
	"gitlab.com/miguelit0/toDoApp/models"
	"gitlab.com/miguelit0/toDoApp/repository"
)

func (r *RouteHandlers) routesToDoApp() []Route {
	return []Route{
		{
			Name:        "Create task",
			Method:      fasthttp.MethodPost,
			Pattern:     "/tasks",
			HandlerFunc: r.CreateTask,
		},
		{
			Name:        "Get task",
			Method:      fasthttp.MethodGet,
			Pattern:     "/task",
			HandlerFunc: r.GetTasksByTaskId,
		},
		{
			Name:        "Update task",
			Method:      fasthttp.MethodPut,
			Pattern:     "/tasks",
			HandlerFunc: r.UpdateTask,
		},
		{
			Name:        "Delete task",
			Method:      fasthttp.MethodDelete,
			Pattern:     "/tasks",
			HandlerFunc: r.DeleteTask,
		},
		{
			Name:        "Get all tasks",
			Method:      fasthttp.MethodGet,
			Pattern:     "/tasks",
			HandlerFunc: r.GetTasks,
		},
	}
}

func (r *RouteHandlers) CreateTask(ctx *fasthttp.RequestCtx) {
	var req models.CreateTaskRequest

	if err := json.Unmarshal(ctx.Request.Body(), &req); err != nil {
		helpers.RespError(ctx, err)
	}

	var task = models.Task{
		Id:          uuid.UUID(uuid.New()).String()[:8],
		Title:       req.Title,
		Description: req.Description,
	}

	err := repository.CreateTask(ctx, task)

	if err != nil {
		helpers.RespError(ctx, err)
	}

	message := "Task created successfully"

	helpers.RespOk(ctx, message)

}

func (r *RouteHandlers) UpdateTask(ctx *fasthttp.RequestCtx) {

	var task models.Task

	if err := json.Unmarshal(ctx.Request.Body(), &task); err != nil {
		helpers.RespError(ctx, err)
	}
	err := repository.UpdateTask(ctx, task, task.Id)

	if err != nil {
		helpers.RespError(ctx, err)
	}

	message := "Task updated successfully"
	helpers.RespOk(ctx, message)

}

func (r *RouteHandlers) DeleteTask(ctx *fasthttp.RequestCtx) {

	taskId := string(ctx.QueryArgs().Peek("id"))
	err := repository.DeleteTask(ctx, taskId)

	if err != nil {
		helpers.RespError(ctx, err)
	}

	message := "Task deleted successfully"
	helpers.RespOk(ctx, message)

}

func (r *RouteHandlers) GetTasks(ctx *fasthttp.RequestCtx) {

	tasks, err := repository.GetTasks(ctx)

	if err != nil {
		helpers.RespError(ctx, err)
	}

	helpers.RespOk(ctx, tasks)

}

func (r *RouteHandlers) GetTasksByTaskId(ctx *fasthttp.RequestCtx) {
	taskId := string(ctx.QueryArgs().Peek("id"))
	task, err := repository.GetTaskByTaskId(ctx, taskId)

	if err != nil {
		helpers.RespError(ctx, err)
	}
	helpers.RespOk(ctx, task)

}
