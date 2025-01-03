package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/constant"
	"backend-golang/modules/task/domain/entity"
	"context"
	"fmt"
	"time"
)

type GetTaskUsecase interface {
	ExecGetTask(ctx context.Context, taskID string) (*entity.Task, error)
}

type getTaskUsecaseImpl struct {
	taskReaderRepository TaskReaderRepository
	taskWriterRepository TaskWriterRepository
}

var _ GetTaskUsecase = (*getTaskUsecaseImpl)(nil)

func NewGetTaskUsecase(
	taskReaderRepository TaskReaderRepository,
	taskWriterRepository TaskWriterRepository,
) GetTaskUsecase {
	return &getTaskUsecaseImpl{
		taskReaderRepository: taskReaderRepository,
		taskWriterRepository: taskWriterRepository,
	}
}

func (uc getTaskUsecaseImpl) ExecGetTask(ctx context.Context, taskID string) (*entity.Task, error) {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	taskEntity, err := uc.taskReaderRepository.FindTaskByCondition(ctx, map[string]interface{}{
		"taskId": taskID,
		"userId": userId,
	})
	if err != nil {
		return nil, constant.ErrorNotFoundTask(err)
	}

	defaultTime := mapper.ParseDate("")
	// Check taskEntity dueTime over current time with timezone
	fmt.Println(time.Now())
	if taskEntity.DueDate != defaultTime && taskEntity.DueDate.Before(time.Now()) && taskEntity.Status != "Completed" && taskEntity.Status != "Expired" {
		taskEntity.Status = "Expired"
		// Update task status
		err1 := uc.taskWriterRepository.UpdateTask(ctx, *taskEntity)
		if err1 != nil {
			return nil, constant.ErrorNotFoundTask(err1)
		}
	}

	return taskEntity, nil
}
