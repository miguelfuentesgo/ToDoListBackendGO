package repository

import (
	"context"

	"gitlab.com/miguelit0/toDoApp/models"
)

type Repository interface {
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTaskById(ctx context.Context, taskId string) (*models.Task, error)
	CreateTask(ctx context.Context, task models.Task) error
	UpdateTask(ctx context.Context, task models.Task, taskId string) error
	DeleteTask(ctx context.Context, taskId string) error
	Close() error
}

var implementation Repository

func NewRepository(repo Repository) {
	implementation = repo
}

func GetTasks(ctx context.Context) ([]models.Task, error) {
	return implementation.GetTasks(ctx)
}

func GetTaskByTaskId(ctx context.Context, taskId string) (*models.Task, error) {
	return implementation.GetTaskById(ctx, taskId)
}

func CreateTask(ctx context.Context, task models.Task) error {
	return implementation.CreateTask(ctx, task)
}

func UpdateTask(ctx context.Context, task models.Task, taskId string) error {
	return implementation.UpdateTask(ctx, task, taskId)
}

func DeleteTask(ctx context.Context, taskId string) error {
	return implementation.DeleteTask(ctx, taskId)
}

func Close() error {
	return implementation.Close()
}
