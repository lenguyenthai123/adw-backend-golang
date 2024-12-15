package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
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
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	taskEntity, err := uc.taskReaderRepository.FindTaskByCondition(ctx, map[string]interface{}{
		"taskId": taskID,
		"userId": userId,
	})
	if err != nil {
		return nil, constant.ErrorNotFoundTask(err)
	}

	return taskEntity, nil
}
