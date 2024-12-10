package usecase

import (
	"backend-golang/modules/task/domain/entity"
	"context"
)

type CreateTaskUsecase interface {
	ExecCreateTask(ctx context.Context, userEntity entity.Task) error
}

type createTaskUsecaseImpl struct {
	writerRepo TaskWriterRepository
}

var _ CreateTaskUsecase = (*createTaskUsecaseImpl)(nil)

func NewCreateTaskUsecase(writerRepo TaskWriterRepository) CreateTaskUsecase {
	return &createTaskUsecaseImpl{
		writerRepo: writerRepo,
	}
}

func (uc createTaskUsecaseImpl) ExecCreateTask(ctx context.Context, taskEntity entity.Task) error {
	err := uc.writerRepo.InsertTask(ctx, taskEntity)
	if err != nil {
		return err
	}

	return nil
}
