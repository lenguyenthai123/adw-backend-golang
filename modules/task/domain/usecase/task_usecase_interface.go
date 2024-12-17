package usecase

import (
	"backend-golang/modules/task/domain/entity"
	"context"
	"time"
)

type TaskReaderRepository interface {
	FindTaskByCondition(ctx context.Context, condition map[string]interface{}) (*entity.Task, error)
	FindTaskListByCondition(ctx context.Context, conditions map[string]interface{}) ([]*entity.Task, error)
	// Version 2
	FindTaskListByRangeTime(ctx context.Context, userId, startTime, endTime string) ([]*entity.Task, error)
}

type TaskWriterRepository interface {
	InsertTask(ctx context.Context, taskEntity entity.Task) error
	InsertTaskList(ctx context.Context, taskEntityList []*entity.Task) error
	UpdateTask(ctx context.Context, taskEntity entity.Task) error
	UpdateTaskList(ctx context.Context, userID string, taskEntityList []*entity.Task) error
	DeleteTask(ctx context.Context, userId int, taskID string) error
	DeleteTaskList(_ context.Context, userId int, taskIDs []string) error
	DeleteTaskInRangeTime(ctx context.Context, userId int, startTime, endTime time.Time) error
}
