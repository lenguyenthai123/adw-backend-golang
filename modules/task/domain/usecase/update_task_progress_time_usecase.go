package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"backend-golang/modules/task/domain/entity"
	"context"
)

type UpdateTaskProgressTimeUsecase interface {
	ExecUpdateTaskProgress(ctx context.Context, taskProgressEntity entity.TaskProgress) error
}

type updateTaskProgressTimeUsecaseImpl struct {
	taskWriterRepo TaskProgressWriterRepository
	taskReaderRepo TaskReaderRepository
}

var _ UpdateTaskProgressTimeUsecase = (*updateTaskProgressTimeUsecaseImpl)(nil)

func NewUpdateTaskProgressTimeUsecase(taskReaderRepo TaskReaderRepository, taskProgressWriterRepo TaskProgressWriterRepository) UpdateTaskProgressTimeUsecase {
	return &updateTaskProgressTimeUsecaseImpl{
		taskReaderRepo: taskReaderRepo,
		taskWriterRepo: taskProgressWriterRepo,
	}
}

func (uc updateTaskProgressTimeUsecaseImpl) ExecUpdateTaskProgress(ctx context.Context, taskProgressEntity entity.TaskProgress) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()

	foundedTask, err := uc.taskReaderRepo.FindTaskByCondition(ctx, map[string]interface{}{
		"taskId": taskProgressEntity.TaskID,
		"userId": userId,
	})
	if foundedTask == nil {
		return constant.ErrorNotFoundTask(err)
	}

	err1 := uc.taskWriterRepo.InsertTaskProgressHistory(ctx, taskProgressEntity)
	if err1 != nil {
		return constant.ErrorNotFoundTask(err1)
	}

	return nil
}
