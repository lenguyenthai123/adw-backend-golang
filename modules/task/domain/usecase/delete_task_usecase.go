package usecase

import (
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
	err := uc.writerRepo.DeleteTask(ctx, taskID)
	if err != nil {
		return err
	}

	return nil
}
