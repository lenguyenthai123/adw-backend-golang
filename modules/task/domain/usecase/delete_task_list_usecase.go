package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"context"
)

type DeleteTaskListUsecase interface {
	ExecDeleteTaskList(ctx context.Context, taskIDList []int) error
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

func (uc deleteTaskListUsecaseImpl) ExecDeleteTaskList(ctx context.Context, taskIDList []int) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	err := uc.writerRepo.DeleteTaskList(ctx, userId, taskIDList)
	if err != nil {
		return constant.ErrrorDeleteTaskListFailed(err)
	}

	return nil
}
