package usecase

import (
	"backend-golang/modules/analytics/domain/entity"
	"context"
)

type GetTaskOfEachStatusUsecase interface {
	ExecGetTaskOfEachStatus(ctx context.Context, startTime, endTime string) (*entity.TaskNumberByStatusEntity, error)
}

type getTaskOfEachStatusUsecaseImpl struct {
	TaskProgressReaderRepository TimeProgressReaderRepo
}

var _ GetTaskOfEachStatusUsecase = (*getRatioTotalTimeSpentUsecaseImpl)(nil)

func NewGetTaskOfEachStatusUsecase(taskProgressReaderRepository TimeProgressReaderRepo) GetTaskOfEachStatusUsecase {
	return &getRatioTotalTimeSpentUsecaseImpl{
		TaskProgressReaderRepository: taskProgressReaderRepository,
	}
}

func (uc getRatioTotalTimeSpentUsecaseImpl) ExecGetTaskOfEachStatus(ctx context.Context, startTime, endTime string) (*entity.TaskNumberByStatusEntity, error) {
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
