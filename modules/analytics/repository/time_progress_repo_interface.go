package repository

import "backend-golang/modules/analytics/domain/entity"

type TimeProgressReaderRepo interface {
	GetTotalTimeSpent(userId int, startTime string, endTime string) (string, error)
	GetTimeSpentDaily(userId int, startTime string, endTime string) (*[]entity.DailyProgressEntity, error)
	GetEachTaskProgress(userId int) (*[]entity.TaskProgressEntity, error)
	GetTaskNumberByStatus(userId int, startTime string, endTime string) (*[]entity.TaskNumberByStatusEntity, error)
}
