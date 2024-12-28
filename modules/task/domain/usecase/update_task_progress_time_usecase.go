package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"backend-golang/modules/task/domain/entity"
	"context"
)

type UpdateTaskProgressTimeUsecase interface {
	ExecUpdateTaskProgressList(ctx context.Context, taskProgressEntityList []*entity.TaskProgress) error
}

type updateTaskProgressTimeUsecaseImpl struct {
	taskWriterRepo TaskProgressWriterRepository
	taskReaderRepo TaskReaderRepository
}

var _ UpdateTaskProgressTimeUsecase = (*updateTaskProgressTimeUsecaseImpl)(nil)

func NewUpdateTaskProgressTimeUsecase(taskReaderRepo TaskReaderRepository, taskProgressWriterRepo TaskProgressWriterRepository) UpdateTaskProgressTimeUsecase {
	return &updateTaskProgressTimeUsecaseImpl{
		taskReaderRepo: taskReaderRepo,
		taskWriterRepo: taskProgressWriterRepo,
	}
}

func (uc updateTaskProgressTimeUsecaseImpl) ExecUpdateTaskProgressList(ctx context.Context, taskProgressEntityList []*entity.TaskProgress) error {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	var taskId int
	if len(taskProgressEntityList) == 0 {
		return constant.ErrorNotFoundTask(nil)
	} else {
		taskId = taskProgressEntityList[0].TaskID
	}

	foundedTask, err := uc.taskReaderRepo.FindTaskByCondition(ctx, map[string]interface{}{
		"taskId": taskId,
		"userId": userId,
	})
	if foundedTask == nil {
		return constant.ErrorNotFoundTask(err)
	}

	err1 := uc.taskWriterRepo.InsertTaskProgressListHistory(ctx, taskProgressEntityList)
	if err1 != nil {
		return constant.ErrorNotFoundTask(err1)
	}

	return nil
}
