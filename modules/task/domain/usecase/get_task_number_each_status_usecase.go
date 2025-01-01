package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"context"
)

type GetTaskNumberEachStatusUsecase interface {
	ExecGetTaskNumberEachStatus(ctx context.Context) (map[string]int, error)
}

type getTaskNumberEachStatusUsecaseImpl struct {
	taskReaderRepository TaskReaderRepository
}

var _ GetTaskNumberEachStatusUsecase = (*getTaskNumberEachStatusUsecaseImpl)(nil)

func NewGetTaskNumberEachStatusUsecase(taskReaderRepository TaskReaderRepository) GetTaskNumberEachStatusUsecase {
	return &getTaskNumberEachStatusUsecaseImpl{
		taskReaderRepository: taskReaderRepository,
	}
}

func (uc getTaskNumberEachStatusUsecaseImpl) ExecGetTaskNumberEachStatus(ctx context.Context) (map[string]int, error) {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	taskNumberEachStatus, err := uc.taskReaderRepository.GetTotalTasksOfEachStatus(ctx, userId)
	if err != nil {
		return nil, constant.ErrorNotFoundTask(err)
	}

	return taskNumberEachStatus, nil
}
