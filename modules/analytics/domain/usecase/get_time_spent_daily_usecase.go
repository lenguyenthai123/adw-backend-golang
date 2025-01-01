package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/analytics/domain/entity"
	"backend-golang/modules/task/constant"
	"context"
)

type GetTimeSpentDailyUsecase interface {
	Execute(ctx context.Context, startTime string, endTime string) (*[]entity.DailyProgressEntity, error)
}

type getTimeSpentDailyUsecaseImpl struct {
	timeProgressReaderRepo TimeProgressReaderRepo
}

var _ GetTimeSpentDailyUsecase = (*getTimeSpentDailyUsecaseImpl)(nil)

func NewGetTimeSpentDailyUsecase(timeProgressReaderRepo TimeProgressReaderRepo) GetTimeSpentDailyUsecase {
	return &getTimeSpentDailyUsecaseImpl{
		timeProgressReaderRepo: timeProgressReaderRepo,
	}
}

func (uc getTimeSpentDailyUsecaseImpl) Execute(ctx context.Context, startTime string, endTime string) (*[]entity.DailyProgressEntity, error) {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	timeSpentDaily, err := uc.timeProgressReaderRepo.GetTimeSpentDaily(userId, startTime, endTime)
	if err != nil {
		return nil, constant.ErrorNotFoundTask(err)
	}

	return timeSpentDaily, nil
}
