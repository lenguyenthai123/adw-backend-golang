package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"context"
)

type DeleteTaskUsecase interface {
	ExecDeleteTask(ctx context.Context, taskID string) error
}

type deleteTaskUsecaseImpl struct {
	writerRepo TaskWriterRepository
}

var _ DeleteTaskUsecase = (*deleteTaskUsecaseImpl)(nil)

func NewDeleteTaskUsecase(writerRepo TaskWriterRepository) DeleteTaskUsecase {
	return &deleteTaskUsecaseImpl{
		writerRepo: writerRepo,
	}
}

func (uc deleteTaskUsecaseImpl) ExecDeleteTask(ctx context.Context, taskID string) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	err := uc.writerRepo.DeleteTask(ctx, userId, taskID)
	if err != nil {
		return constant.ErrorNotFoundTask(err)
	}

	return nil
}
