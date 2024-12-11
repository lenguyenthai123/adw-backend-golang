package usecase

import (
	"backend-golang/modules/task/domain/entity"
	"context"
)

type UpdateTaskUsecase interface {
	ExecUpdateTask(ctx context.Context, taskID string, taskEntity entity.Task) error
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

func (uc updateTaskUsecaseImpl) ExecUpdateTask(ctx context.Context, taskID string, taskEntity entity.Task) error {
	err := uc.taskWriterRepo.UpdateTask(ctx, taskEntity)
	if err != nil {
		return err
	}

	return nil
}
