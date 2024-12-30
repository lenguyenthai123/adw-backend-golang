package usecase

import (
	"backend-golang/modules/analytics/domain/entity"
	"context"
)

type GetUserProgressUsecase interface {
	ExecuteGetProgress(ctx context.Context, startTime string, endTime string) (*entity.TaskProgressEntity, error)
}

type getUserProgressUsecaseImpl struct {
	TimeProgressReaderRepo TimeProgressReaderRepo
}

var _ GetUserProgressUsecase = (*getUserProgressUsecaseImpl)(nil)

func NewGetUserProgressUsecase(timeProgressReaderRepo TimeProgressReaderRepo) GetUserProgressUsecase {
	return &getUserProgressUsecaseImpl{
		TimeProgressReaderRepo: timeProgressReaderRepo,
	}
}

func (uc getUserProgressUsecaseImpl) ExecuteGetProgress(ctx context.Context, startTime, endTime string) (*entity.TaskProgressEntity, error) {

	return nil, nil
}
