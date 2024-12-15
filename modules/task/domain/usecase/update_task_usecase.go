package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"backend-golang/modules/task/domain/entity"
	"context"
)

type UpdateTaskUsecase interface {
	ExecUpdateTask(ctx context.Context, taskEntity entity.Task) error
}

type updateTaskUsecaseImpl struct {
	taskWriterRepo TaskWriterRepository
	taskReaderRepo TaskReaderRepository
}

var _ UpdateTaskUsecase = (*updateTaskUsecaseImpl)(nil)

func NewUpdateTaskUsecase(taskReaderRepo TaskReaderRepository, taskWriterRepo TaskWriterRepository) UpdateTaskUsecase {
	return &updateTaskUsecaseImpl{
		taskReaderRepo: taskReaderRepo,
		taskWriterRepo: taskWriterRepo,
	}
}

func (uc updateTaskUsecaseImpl) ExecUpdateTask(ctx context.Context, taskEntity entity.Task) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	taskEntity.UserID = userId

	foundedTask, err := uc.taskReaderRepo.FindTaskByCondition(ctx, map[string]interface{}{
		"taskId": taskEntity.TaskID,
		"userId": taskEntity.UserID,
	})
	if foundedTask == nil {
		return constant.ErrorNotFoundTask(err)
	}

	err1 := uc.taskWriterRepo.UpdateTask(ctx, taskEntity)
	if err1 != nil {
		return constant.ErrorNotFoundTask(err1)
	}

	return nil
}
