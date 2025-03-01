package repository

import (
	"backend-golang/modules/task/domain/entity"
	database "backend-golang/pkgs/dbs/postgres"
	"context"
	"log/slog"
	"time"
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
func (repo taskWriterRepositoryImpl) InsertTaskList(ctx context.Context, taskEntityList []*entity.Task) error {
	return repo.db.Executor.
		Create(&taskEntityList).Error
}

func (repo taskWriterRepositoryImpl) DeleteTask(_ context.Context, userId int, taskID string) error {
	return repo.db.Executor.
		Where("\"taskId\" = ?", taskID).
		Where("\"userId\" = ?", userId).
		Delete(&entity.Task{}).Error
}

func (repo taskWriterRepositoryImpl) DeleteTaskList(_ context.Context, userId int, taskIDs []string) error {
	return repo.db.Executor.
		Where("\"taskId\" IN (?)", taskIDs).
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

func (repo taskWriterRepositoryImpl) UpdateTaskList(ctx context.Context, userID string, taskEntityList []*entity.Task) error {
	for _, task := range taskEntityList {
		slog.Any("task: ", task)
		slog.Any("task: ", *task)
		slog.Any("(*task).TaskID: ", (*task).TaskID)
		slog.Any("task.TaskID: ", task.TaskID)

		err := repo.db.Executor.Model(&entity.Task{}).
			Where("\"taskId\" = ?", (*task).TaskID).
			Where("\"userId\" = ?", userID).
			Updates(&map[string]interface{}{
				"priority":  (*task).Priority,
				"startDate": (*task).StartDate,
				"dueDate":   (*task).DueDate,
			}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo taskWriterRepositoryImpl) DeleteTaskInRangeTime(ctx context.Context, userId int, startTime, endTime time.Time) error {
	// Perform the delete operation
	query := repo.db.Executor.WithContext(ctx).Model(&entity.Task{})
	result := query.
		Where("\"userId\" = ?", userId).
		Where("((\"dueDate\" >= ? AND \"dueDate\" <= ?) OR (\"startDate\" >= ? AND \"startDate\" <= ?))",
			startTime, endTime, startTime, endTime).
		Delete(&entity.Task{})

	// Check for errors
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo taskWriterRepositoryImpl) UpdateTaskListToExpired(ctx context.Context, taskIdList []int) error {
	// Perform the update operation
	query := repo.db.Executor.WithContext(ctx).Model(&entity.Task{})
	result := query.
		Where("\"taskId\" IN (?)", taskIdList).
		Updates(map[string]interface{}{"status": "Expired"})

	// Check for errors
	if result.Error != nil {
		return result.Error
	}
	return nil
}
