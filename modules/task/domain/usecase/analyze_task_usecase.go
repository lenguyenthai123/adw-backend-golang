package usecase

import (
	"backend-golang/pkgs/log"
	"context"
	"time"
)

type AnalyzeTaskUsecase interface {
	ExecAnalyzeTask(ctx context.Context, startTime, endTime time.Time) error
}

type analyzeTaskUsecaseImpl struct {
	readerRepo TaskReaderRepository
}

var _ AnalyzeTaskUsecase = (*analyzeTaskUsecaseImpl)(nil)

func NewAnalyzeTaskUsecase(readerRepo TaskReaderRepository) AnalyzeTaskUsecase {
	return &analyzeTaskUsecaseImpl{
		readerRepo: readerRepo,
	}
}

func (uc analyzeTaskUsecaseImpl) ExecAnalyzeTask(ctx context.Context, startTime, endTime time.Time) error {

	taskList, err := uc.readerRepo.FindTaskListByConditionV2(ctx, map[string]interface{}{
		"DueDate >=": startTime,
		"DueDate <=": endTime})

	log.JsonLogger.Info("ExecAnalyzeTask.find_task_list_by_time_range", taskList)
	if err != nil {
		return err
	}
	return nil
}
