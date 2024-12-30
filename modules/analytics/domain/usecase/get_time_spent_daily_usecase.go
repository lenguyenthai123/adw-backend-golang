package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/analytics/domain/entity"
	"context"
)

type GetTimeSpentDailyUsecase interface {
	Execute(ctx context.Context, startTime string, endTime string) (*[]entity.DailyProgressEntity, error)
}

type getTimeSpentDailyUsecaseImpl struct {
	TimeProgressReaderRepo TimeProgressReaderRepo
}

var _ GetTimeSpentDailyUsecase = (*getTimeSpentDailyUsecaseImpl)(nil)

func NewGetTimeSpentDailyUsecase(timeProgressReaderRepo TimeProgressReaderRepo) GetTimeSpentDailyUsecase {
	return &getTimeSpentDailyUsecaseImpl{
		TimeProgressReaderRepo: timeProgressReaderRepo,
	}
}

func (uc getTimeSpentDailyUsecaseImpl) Execute(ctx context.Context, startTime string, endTime string) (*[]entity.DailyProgressEntity, error) {
	return uc.TimeProgressReaderRepo.GetTimeSpentDaily(ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt(), startTime, endTime)
}
