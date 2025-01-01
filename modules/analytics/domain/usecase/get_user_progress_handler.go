package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/analytics/constant"
	"backend-golang/modules/analytics/domain/entity"
	"context"
)

type GetUserProgressUsecase interface {
	ExecuteGetProgress(ctx context.Context) (*[]entity.TaskProgressEntity, error)
}

type getUserProgressUsecaseImpl struct {
	timeProgressReaderRepo TimeProgressReaderRepo
}

var _ GetUserProgressUsecase = (*getUserProgressUsecaseImpl)(nil)

func NewGetUserProgressUsecase(timeProgressReaderRepo TimeProgressReaderRepo) GetUserProgressUsecase {
	return &getUserProgressUsecaseImpl{
		timeProgressReaderRepo: timeProgressReaderRepo,
	}
}

func (uc getUserProgressUsecaseImpl) ExecuteGetProgress(ctx context.Context) (*[]entity.TaskProgressEntity, error) {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	taskProgressList, err := uc.timeProgressReaderRepo.GetEachTaskProgress(userId)
	if err != nil {
		return nil, constant.ErrorGetUserProgressFail(err)
	}

	return taskProgressList, nil
}
