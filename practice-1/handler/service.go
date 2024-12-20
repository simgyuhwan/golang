package handler

import (
	"api/entity"
	"context"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTaskService AddTaskService RegisterUserService
type ListTaskService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}
