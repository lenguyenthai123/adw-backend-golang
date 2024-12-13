package usecase

import (
	"backend-golang/modules/task/domain/entity"
	"context"
)

type GetTaskListUsecase interface {
	ExecGetTaskList(ctx context.Context, searchFilter entity.TaskSearchFilterEntity) ([]*entity.Task, error)
}

type getTaskListUsecaseImpl struct {
	taskReaderRepository TaskReaderRepository
}

var _ GetTaskListUsecase = (*getTaskListUsecaseImpl)(nil)

func NewGetTaskListUsecase(taskReaderRepository TaskReaderRepository) GetTaskListUsecase {
	return &getTaskListUsecaseImpl{
		taskReaderRepository: taskReaderRepository,
	}
}

func (uc getTaskListUsecaseImpl) ExecGetTaskList(ctx context.Context, searchFilter entity.TaskSearchFilterEntity) ([]*entity.Task, error) {
	// Build conditions for filtering
	conditions := make(map[string]interface{})

	if searchFilter.UserID != nil {
		conditions["userId"] = *searchFilter.UserID
	}
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
		return nil, err
	}

	return tasks, nil
}
