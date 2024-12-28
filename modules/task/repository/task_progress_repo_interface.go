package repository

import (
	"backend-golang/modules/task/domain/entity"
	"context"
)

type TaskProgressReaderRepository interface {
	FindTaskProgressByCondition(ctx context.Context, condition map[string]interface{}) (*entity.TaskProgress, error)
}

type TaskProgressWriterRepository interface {
	InsertTaskProgressListHistory(ctx context.Context, taskProgressEntityList []*entity.TaskProgress) error
}
