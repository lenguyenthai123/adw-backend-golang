package repository

import (
	"backend-golang/modules/task/domain/entity"
	"context"
)

type TaskReaderRepository interface {
	FindTaskByCondition(ctx context.Context, condition map[string]interface{}) (*entity.Task, error)
	FindTaskListByCondition(ctx context.Context, conditions map[string]interface{}) ([]*entity.Task, error)

	// Version 2
	FindTaskListByConditionV2(ctx context.Context, conditions map[string]interface{}) ([]*entity.Task, error)
}

type TaskWriterRepository interface {
	InsertTask(ctx context.Context, taskEntity entity.Task) error
	UpdateTask(ctx context.Context, taskEntity entity.Task) error
	DeleteTask(ctx context.Context, taskID string) error
}
