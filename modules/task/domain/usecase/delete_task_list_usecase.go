package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"context"
)

type DeleteTaskListUsecase interface {
	ExecDeleteTaskList(ctx context.Context, taskIDs []string) error
}

type deleteTaskListUsecaseImpl struct {
	writerRepo TaskWriterRepository
}

var _ DeleteTaskListUsecase = (*deleteTaskListUsecaseImpl)(nil)

func NewDeleteTaskListUsecase(writerRepo TaskWriterRepository) DeleteTaskListUsecase {
	return &deleteTaskListUsecaseImpl{
		writerRepo: writerRepo,
	}
}

func (uc deleteTaskListUsecaseImpl) ExecDeleteTaskList(ctx context.Context, taskIDs []string) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	err := uc.writerRepo.DeleteTaskList(ctx, userId, taskIDs)
	if err != nil {
		return constant.ErrorNotFoundTask(err)
	}

	return nil
}
