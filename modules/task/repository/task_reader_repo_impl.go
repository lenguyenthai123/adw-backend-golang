package repository

import (
	"backend-golang/modules/task/domain/entity"
	database "backend-golang/pkgs/dbs/postgres"
	"context"
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
