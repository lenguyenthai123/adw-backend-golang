package usecase

import (
	"backend-golang/modules/task/domain/entity"
	"context"
)

type GetTaskUsecase interface {
	ExecGetTask(ctx context.Context, taskID string) (*entity.Task, error)
}

type getTaskUsecaseImpl struct {
	taskReaderRepository TaskReaderRepository
}

var _ GetTaskUsecase = (*getTaskUsecaseImpl)(nil)

func NewGetTaskUsecase(taskReaderRepository TaskReaderRepository) GetTaskUsecase {
	return &getTaskUsecaseImpl{
		taskReaderRepository: taskReaderRepository,
	}
}

func (uc getTaskUsecaseImpl) ExecGetTask(ctx context.Context, taskID string) (*entity.Task, error) {
	taskEntity, err := uc.taskReaderRepository.FindTaskByCondition(ctx, map[string]interface{}{"taskId": taskID})
	if err != nil {
		return nil, err
	}

	return taskEntity, nil
}
