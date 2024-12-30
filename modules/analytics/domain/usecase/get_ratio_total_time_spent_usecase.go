package usecase

import (
	"backend-golang/modules/analytics/domain/entity"
	"context"
)

type GetRatioTotalTimeSpentUsecase interface {
	ExecGetRatioTotalTimeSpent(ctx context.Context, startTime, endTime string) (*entity.TimeSpentEntity, error)
}

type getRatioTotalTimeSpentUsecaseImpl struct {
	TaskProgressReaderRepository TimeProgressReaderRepo
}

var _ GetRatioTotalTimeSpentUsecase = (*getRatioTotalTimeSpentUsecaseImpl)(nil)

func NewGetRatioTotalTimeSpentUsecase(taskProgressReaderRepository TimeProgressReaderRepo) GetRatioTotalTimeSpentUsecase {
	return &getRatioTotalTimeSpentUsecaseImpl{
		TaskProgressReaderRepository: taskProgressReaderRepository,
	}
}

func (uc getRatioTotalTimeSpentUsecaseImpl) ExecGetRatioTotalTimeSpent(ctx context.Context, startTime, endTime string) (*entity.TimeSpentEntity, error) {
	// userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()

	// totalTimeSpent, err := uc.TaskProgressReaderRepository.GetTotalTimeSpent(ctx, userId)
	// if err != nil {
	// 	return nil, constant.ErrorGetTotalTimeSpent(err)
	// }

	// return &entity.TimeSpentEntity{
	// 	TotalTimeSpent: totalTimeSpent,
	// }, nil

	return nil, nil
}
