package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"backend-golang/modules/task/domain/entity"
	"context"
)

type ApplyAnalyzedTaskUsecase interface {
	ExecApplyAnalyzedTask(ctx context.Context, taskEntityList entity.TaskApplyAnalyzedTaskEntity) error
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

func (uc applyAnalyzedTaskUsecaseImpl) ExecApplyAnalyzedTask(ctx context.Context, taskEntityList entity.TaskApplyAnalyzedTaskEntity) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()

	if len(taskEntityList.TaskList) == 0 {
		return nil
	}

	err := uc.writerRepo.DeleteTaskInRangeTime(ctx, userId, taskEntityList.StartTime, taskEntityList.EndTime)
	if err != nil {
		return constant.ErrrorUpdateTaskListFailed(err)
	}
	for _, task := range taskEntityList.TaskList {
		task.UserID = userId
	}
	err = uc.writerRepo.InsertTaskList(ctx, taskEntityList.TaskList)

	return nil
}
