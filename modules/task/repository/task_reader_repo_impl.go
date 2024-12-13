package repository

import (
	"backend-golang/modules/task/domain/entity"
	database "backend-golang/pkgs/dbs/postgres"
	"context"
	"fmt"
)

type taskReaderRepositoryImpl struct {
	db database.Database
}

var _ TaskReaderRepository = (*taskReaderRepositoryImpl)(nil)

func NewTaskReaderRepository(db database.Database) TaskReaderRepository {
	return &taskReaderRepositoryImpl{
		db: db,
	}
}

func (repo taskReaderRepositoryImpl) FindTaskByCondition(_ context.Context,
	condition map[string]interface{}) (*entity.Task, error) {
	var taskEntity entity.Task

	err := repo.db.Executor.Where(condition).First(&taskEntity).Error
	if err != nil {
		return nil, err
	}

	return &taskEntity, nil
}

func (repo taskReaderRepositoryImpl) FindTaskListByCondition(ctx context.Context,
	conditions map[string]interface{}) ([]*entity.Task, error) {
	var tasks []*entity.Task
	query := repo.db.Executor.WithContext(ctx).Model(&entity.Task{})

	// **Filter**
	if userId, ok := conditions["userId"]; ok {
		query = query.Where("\"userId\" = ?", userId)
	}
	if status, ok := conditions["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if priority, ok := conditions["priority"]; ok {
		query = query.Where("priority = ?", priority)
	}

	// **Search**
	if search, ok := conditions["search"]; ok {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		query = query.Where("\"taskName\" ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
	}

	// **Sort**
	if sortBy, ok := conditions["sortBy"]; ok {
		order := "asc"
		if ord, exists := conditions["order"]; exists {
			order = ord.(string)
		}
		query = query.Order(fmt.Sprintf("%s %s", sortBy, order))
	}

	var limitNumber interface{}
	// **Pagination**
	if limit, ok := conditions["limit"]; ok {
		query = query.Limit(limit.(int))
		limitNumber = limit
	}
	if page, ok := conditions["page"]; ok && limitNumber != nil {
		offset := (page.(int) - 1) * limitNumber.(int)
		query = query.Offset(offset)
	}

	// Execute the query
	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
