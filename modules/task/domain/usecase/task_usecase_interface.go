package usecase

import (
	"backend-golang/modules/task/domain/entity"
	"context"
)

type TaskReaderRepository interface {
	FindTaskByCondition(ctx context.Context, condition map[string]interface{}) (*entity.Task, error)
}

type TaskWriterRepository interface {
	InsertTask(ctx context.Context, taskEntity entity.Task) error
	UpdateTask(ctx context.Context, taskEntity entity.Task) error
	DeleteTask(ctx context.Context, taskID string) error
}
