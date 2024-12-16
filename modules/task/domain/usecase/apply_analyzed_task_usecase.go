package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"backend-golang/modules/task/domain/entity"
	"context"
)

type ApplyAnalyzedTaskUsecase interface {
	ExecApplyAnalyzedTask(ctx context.Context, taskEntityList []*entity.Task) error
}

type applyAnalyzedTaskUsecaseImpl struct {
	writerRepo TaskWriterRepository
}

var _ ApplyAnalyzedTaskUsecase = (*applyAnalyzedTaskUsecaseImpl)(nil)

func NewApplyAnalyzedTaskUsecase(
	writerRepo TaskWriterRepository) ApplyAnalyzedTaskUsecase {
	return &applyAnalyzedTaskUsecaseImpl{
		writerRepo: writerRepo,
	}
}

func (uc applyAnalyzedTaskUsecaseImpl) ExecApplyAnalyzedTask(ctx context.Context, taskEntityList []*entity.Task) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	if len(taskEntityList) == 0 {
		return nil
	}
	err := uc.writerRepo.UpdateTaskList(ctx, userId, taskEntityList)
	if err != nil {
		return constant.ErrrorUpdateTaskListFailed(err)
	}
	return nil
}
