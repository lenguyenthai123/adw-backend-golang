package repository

import (
	"backend-golang/modules/task/domain/entity"
	database "backend-golang/pkgs/dbs/postgres"
	"context"
)

type taskProgressWriterRepositoryImpl struct {
	db database.Database
}

var _ TaskProgressWriterRepository = (*taskProgressWriterRepositoryImpl)(nil)

func NewTaskProgressWriterRepository(db database.Database) TaskProgressWriterRepository {
	return &taskProgressWriterRepositoryImpl{
		db: db,
	}
}

func (repo taskProgressWriterRepositoryImpl) InsertTaskProgressListHistory(ctx context.Context, taskProgressEntityList []*entity.TaskProgress) error {
	return repo.db.Executor.
		Create(&taskProgressEntityList).Error
}
