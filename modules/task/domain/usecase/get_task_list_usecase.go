package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/constant"
	"backend-golang/modules/task/domain/entity"
	"context"
	"time"
)

type GetTaskListUsecase interface {
	ExecGetTaskList(ctx context.Context, searchFilter entity.TaskSearchFilterEntity) ([]*entity.Task, error)
}

type getTaskListUsecaseImpl struct {
	taskReaderRepository TaskReaderRepository
	taskWriterRepository TaskWriterRepository
}

var _ GetTaskListUsecase = (*getTaskListUsecaseImpl)(nil)

func NewGetTaskListUsecase(
	taskReaderRepository TaskReaderRepository,
	taskWriterRepository TaskWriterRepository) GetTaskListUsecase {
	return &getTaskListUsecaseImpl{
		taskReaderRepository: taskReaderRepository,
		taskWriterRepository: taskWriterRepository,
	}
}

func (uc getTaskListUsecaseImpl) ExecGetTaskList(ctx context.Context, searchFilter entity.TaskSearchFilterEntity) ([]*entity.Task, error) {
	// Build conditions for filtering
	conditions := make(map[string]interface{})

	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()
	conditions["userId"] = userId

	if searchFilter.Status != nil {
		conditions["status"] = *searchFilter.Status
	}
	if searchFilter.Priority != nil {
		conditions["priority"] = *searchFilter.Priority
	}
	if searchFilter.Search != nil && *searchFilter.Search != "" {
		conditions["search"] = *searchFilter.Search
	}
	if searchFilter.SortBy != nil && searchFilter.Order != nil {
		conditions["sortBy"] = *searchFilter.SortBy
		conditions["order"] = *searchFilter.Order
	}
	if searchFilter.Limit != nil {
		conditions["limit"] = *searchFilter.Limit
	}
	if searchFilter.Page != nil {
		conditions["page"] = *searchFilter.Page
	}

	tasks, err := uc.taskReaderRepository.FindTaskListByCondition(ctx, conditions)
	if err != nil {
		return nil, constant.ErrorNotFoundTaskList(err)
	}

	// Add all task over due date to a list
	tasksOverDueDate := []int{}

	// Update all task over due date
	for _, task := range tasks {
		if task.DueDate.Before(time.Now()) && task.Status != "Completed" && task.Status != "Expired" {
			task.Status = "Expired"
			tasksOverDueDate = append(tasksOverDueDate, task.TaskID)
		}
	}

	if len(tasksOverDueDate) > 0 {
		err1 := uc.taskWriterRepository.UpdateTaskListToExpired(ctx, tasksOverDueDate)
		if err1 != nil {
			return nil, constant.ErrorSystem(err1)
		}
	}

	return tasks, nil
}
