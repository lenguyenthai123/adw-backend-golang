package repository

import (
	"backend-golang/modules/task/domain/entity"
	database "backend-golang/pkgs/dbs/postgres"
	"context"
)

type taskWriterRepositoryImpl struct {
	db database.Database
}

var _ TaskWriterRepository = (*taskWriterRepositoryImpl)(nil)

func NewTaskWriterRepository(db database.Database) TaskWriterRepository {
	return &taskWriterRepositoryImpl{
		db: db,
	}
}

func (repo taskWriterRepositoryImpl) InsertTask(_ context.Context, taskEntity entity.Task) error {
	return repo.db.Executor.
		Create(&taskEntity).Error
}

func (repo taskWriterRepositoryImpl) DeleteTask(_ context.Context, userId int, taskID string) error {
	return repo.db.Executor.
		Where("\"taskId\" = ?", taskID).
		Where("\"userId\" = ?", userId).
		Delete(&entity.Task{}).Error
}

func (repo taskWriterRepositoryImpl) UpdateTask(_ context.Context, taskEntity entity.Task) error {
	return repo.db.Executor.
		Model(&entity.Task{}).
		Where("\"taskId\" = ?", taskEntity.TaskID).
		Where("\"userId\" = ?", taskEntity.UserID).
		Updates(&taskEntity).Error
}
